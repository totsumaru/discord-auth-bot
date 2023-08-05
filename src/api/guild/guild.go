package guild

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"io"
	"net/http"
	"strings"
)

type Res struct {
	Servers []resServer `json:"servers"`
}

type resServer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

// 自分が管理できるサーバーの一覧を取得します
func MyGuilds(e *gin.Engine) {
	e.GET("/api/guild", func(c *gin.Context) {
		header := c.GetHeader(api.HeaderAuthorization)
		discordToken := strings.TrimPrefix(header, "Bearer ")

		if discordToken == "" {
			c.JSON(http.StatusBadRequest, "トークンに値が設定されていません")
			return
		}

		res := Res{
			Servers: []resServer{},
		}

		allGuilds, err := getAllGuilds(discordToken)
		if err != nil {
			c.JSON(http.StatusTooManyRequests, "You are being rate limited.")
			return
		}

		const iconURLTmpl = "https://cdn.discordapp.com/icons/%s/%s.png"

		// botが導入されているサーバーのみをレスポンスとして返します
		s := discord.Session
		guilds := s.State.Guilds

		for _, v := range allGuilds {
			for _, guild := range guilds {
				if v.ID == guild.ID {
					res.Servers = append(res.Servers, resServer{
						ID:      v.ID,
						Name:    v.Name,
						IconURL: fmt.Sprintf(iconURLTmpl, v.ID, v.IconHash),
					})
				}
			}
		}

		c.JSON(http.StatusOK, res)
	})
}

type getAllGuildsRes struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IconHash string `json:"icon"`
}

// 自分が所属しているDiscordサーバーを全て取得します
func getAllGuilds(discordToken string) ([]getAllGuildsRes, error) {
	res := make([]getAllGuildsRes, 0)

	// tokenから参加しているDiscordの一覧を取得
	const url = "https://discordapp.com/api/users/@me/guilds"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, errors.NewError("httpリクエストを作成できません", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", discordToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, errors.NewError("httpリクエストを実行できません", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, errors.NewError("bodyを読み取れません", err)
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return res, errors.NewError("jsonを構造体に変換できません", err)
	}

	return res, nil
}
