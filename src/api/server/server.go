package server

import (
	"github.com/gin-gonic/gin"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/permission"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/verify"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"sort"
)

// レスポンスです
type Res struct {
	Server       res.Server               `json:"server"`
	Roles        []res.RoleWithPermission `json:"roles"`
	Subscriber   res.User                 `json:"subscriber"`
	OperatorRole []res.Role               `json:"operator_role"`
	Status       string                   `json:"status"`
}

// そのサーバーのデフォルトの権限を取得します
func Server(e *gin.Engine) {
	// ?server_id=xxx
	e.GET("/api/server", func(c *gin.Context) {
		authHeader := c.GetHeader(verify.HeaderAuthorization)
		serverID := c.Query("server_id")

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

		apiRes, err := expose.FindByID(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		s := discord.Session

		guild, err := s.Guild(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		roles := guild.Roles

		// ロールをPosition順にソートします
		sort.Slice(roles, func(i, j int) bool {
			return roles[i].Position > roles[j].Position
		})

		subs := res.User{}
		if apiRes.SubscriberID != "" {
			subscriber, err := discord.Session.User(apiRes.SubscriberID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "エラーが発生しました")
				return
			}

			subs.ID = subscriber.ID
			subs.Name = subscriber.Username
			subs.IconURL = subscriber.AvatarURL("")
		}

		allRoles := make([]res.RoleWithPermission, 0)
		for _, role := range roles {
			rr := res.RoleWithPermission{
				Role: res.Role{
					ID:    role.ID,
					Name:  role.Name,
					Color: role.Color,
				},
				Permission: permission.CheckPermission(role.Permissions),
			}
			allRoles = append(allRoles, rr)
		}

		operator := make([]res.Role, 0)
		for _, operatorRoleID := range apiRes.OperatorRoleID {
			for _, role := range roles {
				if operatorRoleID == role.ID {
					resRole := res.Role{
						ID:    role.ID,
						Name:  role.Name,
						Color: role.Color,
					}
					operator = append(operator, resRole)
				}
			}
		}

		r := Res{
			Server: res.Server{
				ID:      guild.ID,
				Name:    guild.Name,
				IconURL: guild.IconURL(""),
			},
			Roles:        allRoles,
			Subscriber:   subs,
			OperatorRole: operator,
			Status:       apiRes.Status,
		}

		c.JSON(http.StatusOK, r)
	})
}
