package apiErr

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"log"
)

// HandleError はエラー応答を一元化して処理します。
func HandleError(c *gin.Context, statusCode int, displayMessage string, err error) {
	logContent := displayMessage
	if err != nil {
		log.Println(displayMessage + err.Error())
	}

	c.JSON(statusCode, gin.H{
		"error": logContent,
	})

	// DiscordにLogを送信します
	errors.SendDiscordLogWithContext(c, errors.NewError(displayMessage, err))

	c.Abort() // これにより、その後のハンドラが実行されないようにします
}
