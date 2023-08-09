package server

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
)

// サーバーの情報のレスポンスです
type Res struct {
	ServerID   string `json:"server_id"`
	Subscriber struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		IconURL string `json:"icon_url"`
	} `json:"subscriber"`
	OperatorRoleID []string `json:"operator_role_id"`
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

		res := Res{
			ServerID:       serverID,
			OperatorRoleID: apiRes.OperatorRoleID,
		}

		if apiRes.SubscriberID != "" {
			subscriber, err := discord.Session.User(apiRes.SubscriberID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "エラーが発生しました")
				return
			}

			res.Subscriber.ID = subscriber.ID
			res.Subscriber.Name = subscriber.Username
			res.Subscriber.IconURL = subscriber.AvatarURL("")
		}

		c.JSON(http.StatusOK, headerRes)
	})
}
