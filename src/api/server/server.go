package server

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/shared"
	"log"
	"net/http"
)

func Server(e *gin.Engine) {
	e.GET("/api/server", server) // ?server_id=xxx
}

func server(c *gin.Context) {
	authHeader := c.GetHeader(shared.HeaderAuthorization)

	discordID, err := shared.GetDiscordIDFromAuthHeader(authHeader)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, discordID)
}
