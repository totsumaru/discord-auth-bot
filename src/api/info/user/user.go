package user

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
)

// レスポンスです
type Res struct {
	User res.User `json:"user"`
}

// ユーザーの情報を取得します
func InfoUser(e *gin.Engine) {
	e.GET("/api/info/user", func(c *gin.Context) {
		authHeader := c.GetHeader(api.HeaderAuthorization)

		apiRes, err := api.GetAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		s := discord.Session

		u, err := s.User(apiRes.DiscordID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		r := Res{
			User: res.User{
				ID:      u.ID,
				Name:    u.Username,
				IconURL: u.AvatarURL(""),
			},
		}

		c.JSON(http.StatusOK, r)
	})
}
