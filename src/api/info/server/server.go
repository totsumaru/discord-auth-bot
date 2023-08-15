package server

import (
	"github.com/gin-gonic/gin"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/res"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/verify"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
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

// リクエストBodyです
type PatchRequestBody struct {
	RoleID []string `json:"role_id"`
}

// オペレーターロールを変更します
func patch(c *gin.Context) {
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

	var reqBody PatchRequestBody
	if err := c.BindJSON(&reqBody); err != nil {
		apiErr.HandleError(c, 500, "リクエストbodyを変換できません", err)
		return
	}

	if err := expose.UpdateOperatorRoleID(serverID, reqBody.RoleID); err != nil {
		apiErr.HandleError(c, 500, "ロール情報を更新できません", err)
		return
	}

	c.JSON(http.StatusOK, "")
}
