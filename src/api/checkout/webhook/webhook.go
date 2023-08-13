package webhook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74/webhook"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"net/http"
	"os"
)

// Stripeからコールされるwebhookです
func Webhook(e *gin.Engine) {
	e.POST("/api/webhook/", func(c *gin.Context) {
		var b []byte
		if err := c.BindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, "不正なリクエストです")
			return
		}

		header := c.GetHeader("Stripe-Signature")
		webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

		event, err := webhook.ConstructEvent(b, header, webhookSecret)
		if err != nil {
			c.JSON(http.StatusBadRequest, "不正なリクエストです")
			return
		}

		// イベントオブジェクトのdocument
		// https://billing.stripe.com/p/login/test_00geXT3KO9jS2CAaEE
		switch event.Type {
		case "checkout.session.completed":
			// Checkout で顧客が「支払う」または「登録」ボタンをクリックすると送信され、新しい購入が通知されます。
			// `CUSTOMER_ID`をDBに登録する必要があります。
			customerID := event.Data.Object["customer"].(string)
			subscriptionID := event.Data.Object["id"].(string)
			metadata := event.Data.Object["metadata"].(map[string]string)
			guildID := metadata["guild_id"]
			discordID := metadata["discord_id"]

			if err = expose.StartSubscription(
				guildID, discordID, customerID, subscriptionID,
			); err != nil {
				c.JSON(http.StatusInternalServerError, "サブスクリプションの開始情報を作成できません")
				return
			}
		case "invoice.paid":
			// TODO: 実装
			// 請求期間ごとに、支払いが成功すると送信されます。
			// ステータスをDBに保存します。
		case "customer.subscription.deleted":
			// 顧客のサブスクリプションが終了すると送信されます。
			metadata := event.Data.Object["metadata"].(map[string]string)
			guildID := metadata["guild_id"]

			if err = expose.DeleteSubscription(guildID); err != nil {
				c.JSON(http.StatusInternalServerError, "サブスクリプションの開始情報を作成できません")
				return
			}
		case "invoice.payment_failed":
			customerID := event.Data.Object["customer"].(string)
			subscriptionID := event.Data.Object["id"].(string)
			metadata := event.Data.Object["metadata"].(map[string]string)
			guildID := metadata["guild_id"]
			discordID := metadata["discord_id"]

			resObj := map[string]string{
				"customerID":     customerID,
				"subscriptionID": subscriptionID,
				"guildID":        guildID,
				"discordID":      discordID,
			}

			errors.SendDiscord(fmt.Errorf(
				"サブスクリプションの支払いに失敗したユーザーがいます。: %v", resObj,
			))
		default:

		}

		c.JSON(http.StatusOK, "")
	})
}
