package user

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
)

type Res struct {
	DiscordID string `json:"discord_id"`
	IconURL   string `json:"icon_url"`
}

// ユーザーの情報を取得します
func InfoUser(e *gin.Engine) {
	e.GET("/api/info/user", func(c *gin.Context) {
		authHeader := c.GetHeader(api.HeaderAuthorization)

		res, err := api.GetAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		s := discord.Session

		u, err := s.User(res.DiscordID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		c.JSON(http.StatusOK, Res{
			DiscordID: u.ID,
			IconURL:   u.AvatarURL(""),
		})
	})
}
