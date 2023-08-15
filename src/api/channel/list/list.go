package list

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/permission"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/verify"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

// レスポンスです
type Res struct {
	Server   res.Server    `json:"server"`
	Channels []res.Channel `json:"channels"`
}

// チャンネルの一覧を取得します
func ChannelList(e *gin.Engine) {
	// ?server_id=xxx
	e.GET("/api/channel/list", func(c *gin.Context) {
		serverID := c.Query("server_id")
		authHeader := c.GetHeader(verify.HeaderAuthorization)

		// verify
		{
			if serverID == "" || authHeader == "" {
				apiErr.HandleError(c, 400, "リクエストが不正です", nil)
				return
			}

			headerRes, err := verify.GetAuthHeader(authHeader)
			if err != nil {
				apiErr.HandleError(c, 401, "トークンの認証に失敗しました", err)
				return
			}

			if err = verify.CanOperate(serverID, headerRes.DiscordID); err != nil {
				apiErr.HandleError(c, 401, "必要な権限を持っていません", err)
				return
			}
		}

		s := discord.Session

		guild, err := s.Guild(serverID)
		if err != nil {
			apiErr.HandleError(c, 500, "サーバー情報を取得できません", err)
			return
		}

		channels, err := s.GuildChannels(serverID)
		if err != nil {
			apiErr.HandleError(c, 500, "チャンネル情報を取得できません", err)
			return
		}

		// ロールをPosition順にソートします
		channels = getSortedGuildChannels(channels)

		r := Res{
			Server: res.Server{
				ID:      guild.ID,
				Name:    guild.Name,
				IconURL: guild.IconURL(""),
			},
			Channels: []res.Channel{},
		}

		for _, ch := range channels {
			resCh := res.Channel{
				ID:   ch.ID,
				Name: ch.Name,
				Type: permission.ConvertChannelType(ch.Type),
			}
			r.Channels = append(r.Channels, resCh)
		}

		c.JSON(http.StatusOK, r)
	})
}

// チャンネルをDiscordでの表示順にソートします
func getSortedGuildChannels(channels []*discordgo.Channel) []*discordgo.Channel {
	// Divide channels into categories and others
	categories := make([]*discordgo.Channel, 0)
	nonCategories := make([]*discordgo.Channel, 0)
	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildCategory {
			categories = append(categories, channel)
		} else {
			nonCategories = append(nonCategories, channel)
		}
	}

	// Sort categories by position
	sort.SliceStable(categories, func(i, j int) bool {
		return categories[i].Position < categories[j].Position
	})

	// Sort non-categories by position
	sort.SliceStable(nonCategories, func(i, j int) bool {
		return nonCategories[i].Position < nonCategories[j].Position
	})

	// Append channels without categories to the sortedChannels
	sortedChannels := make([]*discordgo.Channel, 0)
	for _, channel := range nonCategories {
		if channel.ParentID == "" {
			sortedChannels = append(sortedChannels, channel)
		}
	}

	// For each category, append the category itself and its channels to the sortedChannels
	for _, category := range categories {
		sortedChannels = append(sortedChannels, category)
		categoryChannels := make([]*discordgo.Channel, 0)
		for _, channel := range nonCategories {
			if channel.ParentID == category.ID {
				categoryChannels = append(categoryChannels, channel)
			}
		}
		// Sort category channels by position again
		sort.SliceStable(categoryChannels, func(i, j int) bool {
			return categoryChannels[i].Position < categoryChannels[j].Position
		})
		// Append sorted category channels
		sortedChannels = append(sortedChannels, categoryChannels...)
	}

	return sortedChannels
}
