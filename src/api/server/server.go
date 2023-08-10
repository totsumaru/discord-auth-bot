package server

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/permission"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

// レスポンスです
type Res struct {
	Server res.Server `json:"server"`
	Roles  []res.Role `json:"roles"`
}

// そのサーバーのデフォルトの権限を取得します
func Server(e *gin.Engine) {
	// ?server_id=xxx
	e.GET("/api/server", func(c *gin.Context) {
		//authHeader := c.GetHeader(api.HeaderAuthorization)
		//
		//discordID, err := api.GetDiscordIDFromAuthHeader(authHeader)
		//if err != nil {
		//	log.Fatal(err)
		//}
		serverID := c.Query("server_id")

		s := discord.Session

		guild, err := s.Guild(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		roles := guild.Roles

		// ロールをPosition順にソートします
		sort.Slice(roles, func(i, j int) bool {
			return roles[i].Position > roles[j].Position
		})

		r := Res{
			Server: res.Server{
				ID:      guild.ID,
				Name:    guild.Name,
				IconURL: guild.IconURL(""),
			},
			Roles: []res.Role{},
		}

		for _, role := range roles {
			rr := res.Role{
				ID:         role.ID,
				Name:       role.Name,
				Color:      role.Color,
				Permission: permission.CheckPermission(role.Permissions),
			}
			r.Roles = append(r.Roles, rr)
		}

		c.JSON(http.StatusOK, r)
	})
}
