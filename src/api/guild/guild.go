package guild

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"io"
	"net/http"
	"strings"
)

const iconURLTmpl = "https://cdn.discordapp.com/icons/%s/%s.png"

// レスポンスです
type Res struct {
	Servers []res.Server `json:"servers"`
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

		r := Res{
			Servers: []res.Server{},
		}

		myGuilds, err := getAllGuilds(discordToken)
		if err != nil {
			c.JSON(http.StatusTooManyRequests, "You are being rate limited.")
			return
		}

		me, err := getUser(discordToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "ユーザー情報を取得できません")
			return
		}

		// botが導入されているサーバーのみをレスポンスとして返します
		s := discord.Session
		botGuilds := s.State.Guilds

		for _, myGuild := range myGuilds {
			for _, botGuild := range botGuilds {
				// 参加しているサーバーが一致した場合
				if myGuild.ID == botGuild.ID {
					// owner,adminロール保持,operatorロール保持の場合はOK
					ok, err := api.VerifyUser(myGuild.ID, me.ID)
					if err != nil {
						c.JSON(http.StatusInternalServerError, "認証できません")
						return
					}
					if ok {
						iconUrl := ""
						if myGuild.IconHash != "" {
							iconUrl = fmt.Sprintf(iconURLTmpl, myGuild.ID, myGuild.IconHash)
						}
						r.Servers = append(r.Servers, res.Server{
							ID:      myGuild.ID,
							Name:    myGuild.Name,
							IconURL: iconUrl,
						})
					}
				}
			}
		}

		c.JSON(http.StatusOK, r)
	})
}

// ユーザー情報のレスポンスです
// doc: https://discord.com/developers/docs/resources/user#user-object
type getUserRes struct {
	ID string `json:"id"`
}

// 自分のユーザー情報を取得します
func getUser(discordToken string) (getUserRes, error) {
	r := getUserRes{}

	// tokenから参加しているDiscordの一覧を取得
	const url = "https://discordapp.com/api/users/@me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return r, errors.NewError("httpリクエストを作成できません", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", discordToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return r, errors.NewError("httpリクエストを実行できません", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, errors.NewError("bodyを読み取れません", err)
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return r, errors.NewError("jsonを構造体に変換できません", err)
	}

	return r, nil
}

// 所属しているサーバーのレスポンスです
type getAllGuildsRes struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IconHash string `json:"icon"`
	Owner    bool   `json:"owner"`
}

// 自分が所属しているDiscordサーバーを全て取得します
func getAllGuilds(discordToken string) ([]getAllGuildsRes, error) {
	r := make([]getAllGuildsRes, 0)

	// tokenから参加しているDiscordの一覧を取得
	const url = "https://discordapp.com/api/users/@me/guilds"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return r, errors.NewError("httpリクエストを作成できません", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", discordToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return r, errors.NewError("httpリクエストを実行できません", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, errors.NewError("bodyを読み取れません", err)
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return r, errors.NewError("jsonを構造体に変換できません", err)
	}

	return r, nil
}
