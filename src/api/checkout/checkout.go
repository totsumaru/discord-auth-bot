package checkout

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/techstart35/discord-auth-bot/src/shared/api"
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
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		ds := discord.Session
		guild, err := ds.Guild(serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		u, err := ds.User(headerRes.DiscordID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "エラーが発生しました")
			return
		}

		successURL := "https://google.com"

		params := &stripe.CheckoutSessionParams{
			SuccessURL: &successURL,
			//CancelURL:  "https://example.com/canceled.html",
			Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
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
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, Res{
			RedirectURL: s.URL,
		})
	})
}
