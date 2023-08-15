package portal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/billingportal/session"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/verify"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"net/http"
	"os"
)

// レスポンスです
type Res struct {
	RedirectURL string `json:"redirect_url"`
}

// カスタマーポータルのURLを作成します
func CreateCustomerPortal(e *gin.Engine) {
	// カスタマーポータルのURLを作成します
	e.POST("/api/portal", func(c *gin.Context) {
		stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
		serverID := c.Query("server_id")
		authHeader := c.GetHeader(verify.HeaderAuthorization)

		var userID string

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
			userID = headerRes.DiscordID

			if err = verify.CanOperate(serverID, headerRes.DiscordID); err != nil {
				apiErr.HandleError(c, 401, "必要な権限を持っていません", err)
				return
			}
		}

		// そのサーバーの支払い者が本人であるかを確認
		apiRes, err := expose.FindByID(serverID)
		if err != nil {
			apiErr.HandleError(c, 500, "サーバーIDで情報を取得できません", err)
			return
		}
		if apiRes.SubscriberID != userID {
			apiErr.HandleError(c, 401, "支払い者本人ではありません", err)
			return
		}

		// customerIdからカスタマーポータルURLを作成
		params := &stripe.BillingPortalSessionParams{
			Customer: stripe.String(apiRes.CustomerID),
			ReturnURL: stripe.String(fmt.Sprintf(
				"%s/dashboard/%s/config",
				os.Getenv("FRONTEND_URL"),
				serverID,
			)),
		}

		s, err := session.New(params)
		if err != nil {
			apiErr.HandleError(c, 500, "stripeのsessionを作成できません", err)
			return
		}

		c.JSON(http.StatusOK, Res{
			RedirectURL: s.URL,
		})
	})
}
