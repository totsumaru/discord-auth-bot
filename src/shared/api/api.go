package api

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/golang-jwt/jwt"
	"github.com/techstart35/discord-auth-bot/src/api/permission"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"github.com/techstart35/discord-auth-bot/src/shared/map/seeker"
	"os"
	"strings"
)

const HeaderAuthorization = "Authorization"

// TODO: 有効期限の確認を行う

type Res struct {
	DiscordID string
}

// Header(Bearer xxx)からdiscordIDを取得します
// ヘッダーの取得例) authHeader := c.GetHeader(shared.HeaderAuthorization)
func GetAuthHeader(authHeader string) (Res, error) {
	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	// トークンを解析
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("メソッドが期待した値と異なります: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return Res{}, errors.NewError("認証できません", err)
	}

	// トークンが有効なら、ユーザーはログインしていると判断できる
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return Res{
			DiscordID: seeker.Str(claims, []string{"user_metadata", "provider_id"}),
		}, nil
	} else {
		return Res{}, errors.NewError("トークンが無効です")
	}
}

// そのユーザーがサーバーの操作権限を持っていることを確認します
func VerifyUser(serverID, userID string) (bool, error) {
	s := discord.Session

	member, err := s.GuildMember(serverID, userID)
	if err != nil {
		errors.SendDiscord(err)
		return false, nil
	}

	// サーバーにユーザーが存在している場合、操作ロールの有無を確認します
	if member.User.ID == userID {
		// APIからDBに登録されているサーバー設定を取得
		apiRes, err := expose.FindByID(serverID)
		if err != nil {
			return false, errors.NewError("サーバー情報を取得できません", err)
		}

		// サーバーの全てのロールを取得
		guildRoles, err := s.GuildRoles(serverID)
		if err != nil {
			return false, errors.NewError("サーバーのロール一覧を取得できません", err)
		}

		roleMap := make(map[string]*discordgo.Role)
		for _, role := range guildRoles {
			roleMap[role.ID] = role
		}

		// 操作ロールを持っているか確認します
		for _, roleID := range member.Roles {
			// そのロールが管理者権限を持っている場合はtrueを返します
			if permission.HasPermission(roleMap[roleID].Permissions, discordgo.PermissionAdministrator) {
				return true, nil
			}

			for _, operatorRole := range apiRes.OperatorRoleID {
				if roleID == operatorRole {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
