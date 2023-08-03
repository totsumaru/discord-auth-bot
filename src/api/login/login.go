package login

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/shared"
	"log"
	"net/http"
)

func Login(e *gin.Engine) {
	e.GET("/api/login", login)
}

func login(c *gin.Context) {
	authHeader := c.GetHeader(shared.HeaderAuthorization)

	discordID, err := shared.GetDiscordIDFromAuthHeader(authHeader)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, discordID)
}
