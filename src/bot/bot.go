package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

// botが追加された時のハンドラーです
func GuildCreateHandler(s *discordgo.Session, m *discordgo.GuildCreate) {
	// TODO: コマンドを登録します

	// サーバーの情報をDBに保存します
	errors.SendDiscord(fmt.Errorf("テストです"))
	if err := expose.CreateServer(m.Guild.ID); err != nil {
		errors.SendDiscord(err)
		return
	}
}
