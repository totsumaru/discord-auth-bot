package channel

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/shared"
	"log"
	"net/http"
)

func Channel(e *gin.Engine) {
	e.GET("/api/channel", channel) // ?server_id=xxx&channel_id=xxx
}

func channel(c *gin.Context) {
	authHeader := c.GetHeader(shared.HeaderAuthorization)

	discordID, err := shared.GetDiscordIDFromAuthHeader(authHeader)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, discordID)
}
