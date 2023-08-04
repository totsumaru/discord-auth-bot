package server

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/permission"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

type Res struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	ID         string                    `json:"id"`
	Name       string                    `json:"name"`
	Permission permission.RolePermission `json:"permission"`
}

func Server(e *gin.Engine) {
	e.GET("/api/server", server) // ?server_id=xxx
}

func server(c *gin.Context) {
	//authHeader := c.GetHeader(api.HeaderAuthorization)
	//
	//discordID, err := api.GetDiscordIDFromAuthHeader(authHeader)
	//if err != nil {
	//	log.Fatal(err)
	//}
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

	res := Res{}
	for _, role := range roles {
		r := Role{
			ID:         role.ID,
			Name:       role.Name,
			Permission: permission.CheckPermission(role.Permissions),
		}
		res.Roles = append(res.Roles, r)
	}

	c.JSON(http.StatusOK, res)
}
