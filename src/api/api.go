package api

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/channel"
	channelList "github.com/techstart35/discord-auth-bot/src/api/channel/list"
	"github.com/techstart35/discord-auth-bot/src/api/guild"
	infoServer "github.com/techstart35/discord-auth-bot/src/api/info/server"
	"github.com/techstart35/discord-auth-bot/src/api/info/user"
	"github.com/techstart35/discord-auth-bot/src/api/server"
)

// ルートを設定します
func RegisterRouter(e *gin.Engine) {
	Route(e)
	server.Server(e)
	channel.Channel(e)
	channelList.ChannelList(e)
	guild.MyGuilds(e)
	user.InfoUser(e)
	infoServer.InfoServer(e)
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
