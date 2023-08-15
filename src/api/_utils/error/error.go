package apiErr

import (
	"github.com/gin-gonic/gin"
	"log"
)

// HandleError はエラー応答を一元化して処理します。
func HandleError(c *gin.Context, statusCode int, displayMessage string, err error) {
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(statusCode, gin.H{
		"error": displayMessage,
	})
	c.Abort() // これにより、その後のハンドラが実行されないようにします
}
