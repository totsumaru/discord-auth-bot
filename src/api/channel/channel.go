package channel

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/permission"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"log"
	"net/http"
	"sort"
)

type Res struct {
	ChannelID   string    `json:"channel_id"`
	ChannelName string    `json:"channel_name"`
	ChannelType string    `json:"channel_type"`
	IsPrivate   bool      `json:"is_private"`
	Roles       []roleRes `json:"roles"`
}

type roleRes struct {
	ID         string                `json:"id"`
	Name       string                `json:"name"`
	Comment    string                `json:"comment"`    // 推奨設定のコメント(任意)
	Permission permission.Permission `json:"permission"` // チャンネルタイプごとに中身は変更
}

func Channel(e *gin.Engine) {
	e.GET("/api/channel", channel) // ?server_id=xxx&channel_id=xxx
}

func channel(c *gin.Context) {
	//authHeader := c.GetHeader(api.HeaderAuthorization)
	//
	//discordID, err := api.GetDiscordIDFromAuthHeader(authHeader)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, "エラーが発生しました")
	//	return
	//}
	serverID := c.Query("server_id")
	channelID := c.Query("channel_id")

	s := discord.Session
	roles, err := s.GuildRoles(serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "エラーが発生しました")
		return
	}

	// ロールをPosition順にソートします
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Position > roles[j].Position
	})

	ch, err := s.Channel(channelID)
	if err != nil {
		log.Fatal(err)
	}

	isPrivate := isPrivateChannel(ch, serverID)

	res := Res{}
	res.ChannelID = channelID
	res.ChannelName = ch.Name
	res.ChannelType = switchChannelType(ch.Type)
	res.IsPrivate = isPrivate

	for _, role := range roles {
		var isOverrideRole bool

		resRole := roleRes{
			ID:   role.ID,
			Name: role.Name,
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

		if isPrivate &&
			isOverrideRole &&
			rolePm.ViewChannels == false &&
			role.ID != serverID {

			// privateでチャンネルを見るがOFFになっているロールは無駄です
			resRole.Comment = "@everyoneの「チャンネルを見る」をOFFにしたことでプライベートチャンネルになっているため、このロールは設定する必要ありません。"
		}

		if isPrivate && !isOverrideRole {
			// privateチャンネルかつ、上書きされていないロールは、レスポンスに含めません
			continue
		}

		// RolePermission -> チャンネルTypeに応じた型 に型キャスト
		resRole.Permission, err = permission.CastRolePermissionToPermission(rolePm, ch.Type)
		if err != nil {
			errors.SendDiscord(errors.NewError("Permissionの型を変換できません", err))
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		res.Roles = append(res.Roles, resRole)
	}

	c.JSON(http.StatusOK, res)
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
	for _, overRole := range ch.PermissionOverwrites {
		if overRole.ID == serverID {
			return overRole.Deny&discordgo.PermissionViewChannel != 0
		}
	}

	return false
}
