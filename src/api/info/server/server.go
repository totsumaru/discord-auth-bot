package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"net/http"
)

const path = "/api/info/server"

// サーバーの情報のレスポンスです
//
// ロールのPermissionは返しません。
type Res struct {
	Server       res.Server `json:"server"`
	Subscriber   res.User   `json:"subscriber"`
	OperatorRole []res.Role `json:"operator_role"`
}

// サーバーの設定を取得します
func InfoServer(e *gin.Engine) {
	e.PATCH(path, patch)
}

type PatchRequestBody struct {
	RoleID []string `json:"role_id"`
}

// オペレーターロールを変更します
func patch(c *gin.Context) {
	authHeader := c.GetHeader(api.HeaderAuthorization)
	serverID := c.Query("server_id")

	headerRes, err := api.GetAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "エラーが発生しました")
		return
	}

	// ユーザーがサーバーの情報にアクセスできるか検証
	ok, err := api.VerifyUser(serverID, headerRes.DiscordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "エラーが発生しました")
		return
	}
	if !ok {
		c.JSON(http.StatusUnauthorized, "サーバーの情報にアクセスできません")
		return
	}

	var reqBody PatchRequestBody
	if err = c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	if err = expose.UpdateOperatorRoleID(serverID, reqBody.RoleID); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "エラーが発生しました")
		return
	}

	c.JSON(http.StatusOK, "")
}
