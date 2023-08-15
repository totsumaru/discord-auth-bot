package errors

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"log"
)

const (
	ErrorLogChannelID = "1136568705871007875"
	TotsumaruID       = "960104306151948328"
)

// DiscordにLogを送信します
func SendDiscordLogWithContext(c *gin.Context, err error) {
	s := discord.Session

	embed := &discordgo.MessageEmbed{
		Title:       c.Request.URL.String(),
		Description: err.Error(),
	}

	_, err = s.ChannelMessageSendEmbed(ErrorLogChannelID, embed)
	if err != nil {
		log.Println(err)
	}
}

// Discordにメンション付きのアラートLogを送信します
func SendDiscordAlertWithContext(c *gin.Context, err error) {
	s := discord.Session

	embed := &discordgo.MessageEmbed{
		Title:       c.Request.URL.String(),
		Description: err.Error(),
	}

	data := &discordgo.MessageSend{
		Content: fmt.Sprintf("<@%s>", TotsumaruID),
		Embed:   embed,
	}

	_, err = s.ChannelMessageSendComplex(ErrorLogChannelID, data)
	if err != nil {
		log.Println(err)
	}
}

// DiscordにLogを送信します
func SendDiscordLog(err error) {
	s := discord.Session

	embed := &discordgo.MessageEmbed{
		Title:       "API以外でエラーが発生しました",
		Description: err.Error(),
	}

	_, err = s.ChannelMessageSendEmbed(ErrorLogChannelID, embed)
	if err != nil {
		log.Println(err)
	}
}
