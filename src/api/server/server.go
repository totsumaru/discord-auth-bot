package server

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/permission"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

type Res struct {
	ServerName    string `json:"server_name"`
	ServerIconURL string `json:"server_icon_url"`
	Roles         []Role `json:"roles"`
}

type Role struct {
	ID         string                    `json:"id"`
	Name       string                    `json:"name"`
	Color      int                       `json:"color"`
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

	res := Res{
		ServerName:    guild.Name,
		ServerIconURL: guild.IconURL(""),
	}
	for _, role := range roles {
		r := Role{
			ID:         role.ID,
			Name:       role.Name,
			Color:      role.Color,
			Permission: permission.CheckPermission(role.Permissions),
		}
		res.Roles = append(res.Roles, r)
	}

	c.JSON(http.StatusOK, res)
}
