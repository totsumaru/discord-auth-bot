package list

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/permission"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

type Res struct {
	Channels []resChannel `json:"channels"`
}

type resChannel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// チャンネルの一覧を取得します
func ChannelList(e *gin.Engine) {
	// ?server_id=xxx
	e.GET("/api/channel/list", func(c *gin.Context) {
		serverID := c.Query("server_id")

		s := discord.Session

		channels, err := s.GuildChannels(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		// ロールをPosition順にソートします
		channels = getSortedGuildChannels(channels)

		res := Res{
			Channels: []resChannel{},
		}

		for _, ch := range channels {
			resCh := resChannel{
				ID:   ch.ID,
				Name: ch.Name,
				Type: permission.ConvertChannelType(ch.Type),
			}

			res.Channels = append(res.Channels, resCh)
		}

		c.JSON(http.StatusOK, res)
	})
}

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
