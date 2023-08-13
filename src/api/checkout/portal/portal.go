package portal

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/billingportal/session"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
	"net/http"
	"os"
)

// レスポンスです
type Res struct {
	RedirectURL string `json:"redirect_url"`
}

func CreateCustomerPortal(e *gin.Engine) {
	// カスタマーポータルのURLを作成します
	e.POST("/api/portal", func(c *gin.Context) {
		stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
		authHeader := c.GetHeader(api.HeaderAuthorization)
		serverID := c.Query("server_id")

		// discordIDをTokenから取得
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
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		userID := headerRes.DiscordID

		// そのサーバーの支払い者が本人であるかを確認
		apiRes, err := expose.FindByID(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		if apiRes.SubscriberID != userID {
			c.JSON(http.StatusUnauthorized, "支払い者ではありません")
			return
		}

		// customerIdからカスタマーポータルURLを作成
		params := &stripe.BillingPortalSessionParams{
			Customer:  stripe.String(apiRes.CustomerID),
			ReturnURL: stripe.String("https://google.com"),
		}

		s, err := session.New(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		c.JSON(http.StatusOK, Res{
			RedirectURL: s.URL,
		})
	})
}
