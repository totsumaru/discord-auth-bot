package channel

import (
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/permission"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"log"
	"net/http"
	"sort"
)

type Res struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Roles       []Role `json:"roles"`
}

type Role struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Permission permission.Permissions `json:"permission"`
}

func Channel(e *gin.Engine) {
	e.GET("/api/channel", channel) // ?server_id=xxx&channel_id=xxx
}

func channel(c *gin.Context) {
	//authHeader := c.GetHeader(api.HeaderAuthorization)
	//
	//discordID, err := api.GetDiscordIDFromAuthHeader(authHeader)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, "エラーが発生しました")
	//	return
	//}
	serverID := c.Query("server_id")
	channelID := c.Query("channel_id")

	s := discord.Session
	roles, err := s.GuildRoles(serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "エラーが発生しました")
		return
	}

	// ロールをPosition順にソートします
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Position > roles[j].Position
	})

	ch, err := s.Channel(channelID)
	if err != nil {
		log.Fatal(err)
	}

	res := Res{}
	res.ChannelID = channelID
	res.ChannelName = ch.Name
	for _, role := range roles {
		r := Role{
			ID:         role.ID,
			Name:       role.Name,
			Permission: permission.CheckPermission(role.Permissions),
		}
		for _, overRole := range ch.PermissionOverwrites {
			// 上書きロールがある場合は、ここで上書きを実行する
			if role.ID == overRole.ID {
				r.Permission = permission.OverridePermission(r.Permission, overRole.Allow, true)
				r.Permission = permission.OverridePermission(r.Permission, overRole.Deny, false)
			}
		}
		res.Roles = append(res.Roles, r)
	}

	c.JSON(http.StatusOK, res)
}
