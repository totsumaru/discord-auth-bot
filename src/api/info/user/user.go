package user

import (
	"github.com/gin-gonic/gin"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/verify"
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
		authHeader := c.GetHeader(verify.HeaderAuthorization)

		headerRes, err := verify.GetAuthHeader(authHeader)
		if err != nil {
			apiErr.HandleError(c, 401, "トークンの認証に失敗しました", err)
			return
		}

		s := discord.Session

		u, err := s.User(headerRes.DiscordID)
		if err != nil {
			apiErr.HandleError(c, 500, "ユーザー情報を取得できません", err)
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
