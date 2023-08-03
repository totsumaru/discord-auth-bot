package server

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/permission"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"log"
	"net/http"
)

type Res struct {
	Roles []Role
}

type Role struct {
	ID          string
	Name        string
	Color       string
	ChannelView permission.Permission // テスト用
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
		log.Fatal(err)
	}

	for _, role := range roles {
		if role.ID == "998800967665459240" {
			pm := permission.CheckPermission(role)
			c.JSON(http.StatusOK, pm)
			return
		}
	}

	c.JSON(http.StatusOK, Role{})
}
