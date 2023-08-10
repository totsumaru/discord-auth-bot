package list

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

// レスポンスです
//
// Permissionは返しません。
type Res struct {
	Roles []res.Role `json:"roles"`
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

		r := Res{
			Roles: []res.Role{},
		}

		for _, role := range roles {
			rr := res.Role{
				ID:    role.ID,
				Name:  role.Name,
				Color: role.Color,
			}

			r.Roles = append(r.Roles, rr)
		}

		c.JSON(http.StatusOK, r)
	})
}
