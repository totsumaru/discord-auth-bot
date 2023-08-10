package server

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
)

// サーバーの情報のレスポンスです
//
// ロールのPermissionは返しません。
type Res struct {
	Server       res.Server `json:"server"`
	Subscriber   res.User   `json:"subscriber"`
	OperatorRole []res.Role `json:"operator_role"`
}

// サーバーの情報を取得します
func InfoServer(e *gin.Engine) {
	e.GET("/api/info/server", func(c *gin.Context) {
		authHeader := c.GetHeader(api.HeaderAuthorization)
		serverID := c.Query("server_id")

		headerRes, err := api.GetAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		// ユーザーがサーバーの情報にアクセスできるか検証
		ok, err := api.VerifyUser(serverID, headerRes.DiscordID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}
		if !ok {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		apiRes, err := expose.FindByID(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		s := discord.Session
		guild, err := s.Guild(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		subs := res.User{}
		if apiRes.SubscriberID != "" {
			subscriber, err := discord.Session.User(apiRes.SubscriberID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "エラーが発生しました")
				return
			}

			subs.ID = subscriber.ID
			subs.Name = subscriber.Username
			subs.IconURL = subscriber.AvatarURL("")
		}

		operator := make([]res.Role, 0)
		roles, err := s.GuildRoles(guild.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		for _, operatorRoleID := range apiRes.OperatorRoleID {
			for _, role := range roles {
				if operatorRoleID == role.ID {
					resRole := res.Role{
						ID:    role.ID,
						Name:  role.Name,
						Color: role.Color,
					}
					operator = append(operator, resRole)
				}
			}
		}

		r := Res{
			Server: res.Server{
				ID:      serverID,
				Name:    guild.Name,
				IconURL: guild.IconURL(""),
			},
			Subscriber:   subs,
			OperatorRole: []res.Role{},
		}

		c.JSON(http.StatusOK, r)
	})
}
