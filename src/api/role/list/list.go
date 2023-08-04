package list

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

type Res struct {
	Roles []resRole `json:"roles"`
}

type resRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TODO: 認証を追加

// ロールの一覧を取得します
func RoleList(e *gin.Engine) {
	// ?server_id=xxx
	e.GET("/api/role/list", func(c *gin.Context) {
		serverID := c.Query("server_id")

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

		res := Res{
			Roles: []resRole{},
		}

		for _, role := range roles {
			rr := resRole{
				ID:   role.ID,
				Name: role.Name,
			}

			res.Roles = append(res.Roles, rr)
		}

		c.JSON(http.StatusOK, res)
	})
}
