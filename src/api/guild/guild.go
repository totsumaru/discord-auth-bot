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

		res := Res{
			Servers: []resServer{},
		}

		err := func() error {
			header := c.GetHeader(api.HeaderAuthorization)
			discordToken := strings.TrimPrefix(header, "Bearer ")

			// tokenから参加しているDiscordの一覧を取得
			const url = "https://discordapp.com/api/users/@me/guilds"
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return errors.NewError("httpリクエストを作成できません", err)
			}

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", discordToken))

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return errors.NewError("httpリクエストを実行できません", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return errors.NewError("bodyを読み取れません", err)
			}

			type apiResServer struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				IconHash string `json:"icon"`
			}

			var apiRes = make([]apiResServer, 0)
			if err = json.Unmarshal(body, &apiRes); err != nil {
				return errors.NewError("jsonを構造体に変換できません", err)
			}

			const iconURLTmpl = "https://cdn.discordapp.com/icons/%s/%s.png"

			// botが導入されているサーバーのみをレスポンスとして返します
			s := discord.Session
			guilds := s.State.Guilds

			for _, v := range apiRes {
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

			return nil
		}()
		if err != nil {
			errors.SendDiscord(err)
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
