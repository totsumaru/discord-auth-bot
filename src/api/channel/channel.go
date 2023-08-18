package channel

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/permission"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/verify"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

// レスポンスです
type Res struct {
	Server    res.Server               `json:"server"`
	Channel   res.Channel              `json:"channel"`
	IsPrivate bool                     `json:"is_private"`
	Roles     []res.RoleWithPermission `json:"roles"`
	IsActive  bool                     `json:"is_active"`
}

// チャンネルの権限を取得します
//
// - そのサーバーの操作権限が必要です
func Channel(e *gin.Engine) {
	// ?server_id=xxx&channel_id=xxx
	e.GET("/api/channel", func(c *gin.Context) {
		serverID := c.Query("server_id")
		channelID := c.Query("channel_id")
		authHeader := c.GetHeader(verify.HeaderAuthorization)

		// verify
		{
			if serverID == "" || channelID == "" || authHeader == "" {
				apiErr.HandleError(c, 400, "リクエストが不正です", nil)
				return
			}

			headerRes, err := verify.GetAuthHeader(authHeader)
			if err != nil {
				apiErr.HandleError(c, 401, "トークンの認証に失敗しました", err)
				return
			}

			if err = verify.CanOperate(serverID, headerRes.DiscordID); err != nil {
				apiErr.HandleError(c, 401, "必要な権限を持っていません", err)
				return
			}
		}

		s := discord.Session
		guild, err := s.Guild(serverID)
		if err != nil {
			apiErr.HandleError(c, 500, "サーバー情報を取得できません", err)
			return
		}

		ch, err := s.Channel(channelID)
		if err != nil {
			apiErr.HandleError(c, 500, "チャンネル情報を取得できません", err)
			return
		}

		// そのサーバーがProプラン&activeかどうかを判定します
		{
			r, err := expose.FindByID(serverID)
			if err != nil {
				apiErr.HandleError(c, 500, "サーバー情報を取得できません", err)
				return
			}

			// activeでなかったら、ここで終了
			if r.Status != "active" {
				rr := Res{
					Server: res.Server{
						ID:      guild.ID,
						Name:    guild.Name,
						IconURL: guild.IconURL(""),
					},
					Channel: res.Channel{
						ID:   channelID,
						Name: ch.Name,
						Type: switchChannelType(ch.Type),
					},
					IsPrivate: false,
					Roles:     []res.RoleWithPermission{},
					IsActive:  false,
				}

				c.JSON(http.StatusOK, rr)
				return
			}
		}

		roles := guild.Roles

		// ロールをPosition順にソートします
		sort.Slice(roles, func(i, j int) bool {
			return roles[i].Position > roles[j].Position
		})

		isPrivate := isPrivateChannel(ch, serverID)

		resRoles := make([]res.RoleWithPermission, 0)

		for _, role := range roles {
			var isOverrideRole bool

			resRole := res.RoleWithPermission{
				Role: res.Role{
					ID:    role.ID,
					Name:  role.Name,
					Color: role.Color,
				},
			}

			rolePm := permission.CheckPermission(role.Permissions)

			// 上書きロールがある場合は、ここで上書きを実行する
			for _, overRole := range ch.PermissionOverwrites {
				if role.ID == overRole.ID {
					rolePm = permission.OverridePermission(rolePm, overRole.Allow, true)
					rolePm = permission.OverridePermission(rolePm, overRole.Deny, false)
					isOverrideRole = true
				}
			}

			// プライベートチャンネルでViewがOFFになっているロール(管理者を除く)は上書きしている意味がないため、
			// コメントで不要メッセージを追加します。
			if isPrivate &&
				isOverrideRole &&
				rolePm.ViewChannels == false &&
				rolePm.Administrator == false &&
				role.ID != serverID {

				// privateでチャンネルを見るがOFFになっているロールは無駄です
				resRole.Comment = "このロールは設定する必要ありません。@everyoneの「チャンネルを見る」をOFFにしたことで既にプライベートチャンネルになっていますが、二重で「チャンネルを見る」をOFFにしています。"
			}

			// privateチャンネルかつ、上書きされていないロールは、レスポンスに含めません
			// ただし、管理者ロールはレスポンスに含めます。
			if isPrivate && !isOverrideRole && !permission.HasPermission(role.Permissions, discordgo.PermissionAdministrator) {
				continue
			}

			// RolePermission -> チャンネルTypeに応じた型 に型キャスト
			resRole.Permission, err = permission.CastRolePermissionToPermission(rolePm, ch.Type)
			if err != nil {
				apiErr.HandleError(c, 500, "Permissionを変換できません", err)
				return
			}

			resRoles = append(resRoles, resRole)
		}

		r := Res{
			Server: res.Server{
				ID:      guild.ID,
				Name:    guild.Name,
				IconURL: guild.IconURL(""),
			},
			Channel: res.Channel{
				ID:   channelID,
				Name: ch.Name,
				Type: switchChannelType(ch.Type),
			},
			IsPrivate: isPrivate,
			Roles:     resRoles,
			IsActive:  true,
		}

		c.JSON(http.StatusOK, r)
	})
}

// チャンネルタイプを変換します
func switchChannelType(before discordgo.ChannelType) string {
	switch before {
	case discordgo.ChannelTypeGuildText:
		return permission.ChannelTypeText
	case discordgo.ChannelTypeGuildCategory:
		return permission.ChannelTypeCategory
	case discordgo.ChannelTypeGuildNews:
		return permission.ChannelTypeAnnounce
	case discordgo.ChannelTypeGuildForum:
		return permission.ChannelTypeForum
	case discordgo.ChannelTypeGuildVoice:
		return permission.ChannelTypeVC
	case discordgo.ChannelTypeGuildStageVoice:
		return permission.ChannelTypeStage
	}

	return ""
}

// チャンネルがプライベートか判定します
func isPrivateChannel(ch *discordgo.Channel, serverID string) bool {
	for _, overwritePermissionRole := range ch.PermissionOverwrites {
		// @everyoneのView権限がOFFになっていたらPrivateのため、trueを返す
		if overwritePermissionRole.ID == serverID {
			return overwritePermissionRole.Deny&discordgo.PermissionViewChannel != 0
		}
	}

	return false
}
