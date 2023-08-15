package checkout

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/api/_utils/verify"
	"github.com/techstart35/discord-auth-bot/src/shared/discord"
	"net/http"
	"os"
)

// レスポンスです
type Res struct {
	RedirectURL string `json:"redirect_url"`
}

func Checkout(e *gin.Engine) {
	// チェックアウトを作成します
	e.POST("/api/checkout", func(c *gin.Context) {
		stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
		priceId := os.Getenv("STRIPE_PRO_PRICE_ID")

		authHeader := c.GetHeader(verify.HeaderAuthorization)
		serverID := c.Query("server_id")

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

		ds := discord.Session
		guild, err := ds.Guild(serverID)
		if err != nil {
			apiErr.HandleError(c, 500, "サーバー情報を取得できません", err)
			return
		}

		u, err := ds.User(userID)
		if err != nil {
			apiErr.HandleError(c, 500, "ユーザー情報を取得できません", err)
			return
		}

		FEURL := os.Getenv("FRONTEND_URL")

		successURL := fmt.Sprintf("%s/dashboard/%s/success", FEURL, serverID)
		cancelURL := fmt.Sprintf("%s/dashboard/%s/config", FEURL, serverID)

		params := &stripe.CheckoutSessionParams{
			SuccessURL: stripe.String(successURL),
			CancelURL:  stripe.String(cancelURL),
			Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					Price:    stripe.String(priceId),
					Quantity: stripe.Int64(1),
				},
			},
		}
		params.AddMetadata("guild_id", guild.ID)
		params.AddMetadata("discord_id", u.ID)

		s, err := session.New(params)
		if err != nil {
			apiErr.HandleError(c, 500, "stripeのセッションを作成できません", err)
			return
		}

		c.JSON(http.StatusOK, Res{
			RedirectURL: s.URL,
		})
	})
}
