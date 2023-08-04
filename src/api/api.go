package api

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/channel"
	channelList "github.com/techstart35/discord-auth-bot/src/api/channel/list"
	roleList "github.com/techstart35/discord-auth-bot/src/api/role/list"
	"github.com/techstart35/discord-auth-bot/src/api/server"
)

// ルートを設定します
func RegisterRouter(e *gin.Engine) {
	Route(e)
	server.Server(e)
	channel.Channel(e)
	channelList.ChannelList(e)
	roleList.RoleList(e)
}

// ルートです
//
// Note: この関数は削除しても問題ありません
func Route(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.Header("hello", "world")
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
}
