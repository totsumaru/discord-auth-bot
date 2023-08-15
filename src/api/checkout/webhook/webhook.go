package webhook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74/webhook"
	apiErr "github.com/techstart35/discord-auth-bot/src/api/_utils/error"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"io"
	"net/http"
	"os"
)

// Stripeからコールされるwebhookです
func Webhook(e *gin.Engine) {
	e.POST("/api/webhook", func(c *gin.Context) {
		header := c.GetHeader("Stripe-Signature")
		webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			apiErr.HandleError(c, 500, "bodyを読み取れません", err)
			return
		}

		// verify
		event, err := webhook.ConstructEvent(body, header, webhookSecret)
		if err != nil {
			apiErr.HandleError(c, 400, "リクエストが不正です", err)
			return
		}

		// イベントオブジェクトのdocument
		// https://billing.stripe.com/p/login/test_00geXT3KO9jS2CAaEE
		switch event.Type {
		case "checkout.session.completed":
			// Checkout で顧客が「支払う」または「登録」ボタンをクリックすると送信され、新しい購入が通知されます。
			// `CUSTOMER_ID`をDBに登録する必要があります。
			customerID := event.Data.Object["customer"].(string)
			subscriptionID := event.Data.Object["subscription"].(string)
			metadata := event.Data.Object["metadata"].(map[string]interface{})
			guildID := metadata["guild_id"].(string)
			discordID := metadata["discord_id"].(string)

			if err = expose.StartSubscription(
				guildID, discordID, customerID, subscriptionID,
			); err != nil {
				apiErr.HandleError(c, 500, "サブスクリプションを開始できません", err)
				return
			}
		case "invoice.paid":
			// MEMO: 何を実装すべきか不明
			// 請求期間ごとに、支払いが成功すると送信されるイベントです。
			// 初回の支払い時にもコールされます。
		case "customer.subscription.deleted":
			// 顧客のサブスクリプションが終了すると送信されます。
			metadata := event.Data.Object["metadata"].(map[string]string)
			guildID := metadata["guild_id"]

			if err = expose.DeleteSubscription(guildID); err != nil {
				apiErr.HandleError(c, 500, "サブスクリプションを削除できません", err)
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

			errors.SendDiscordAlert(fmt.Errorf(
				"サブスクリプションの支払いに失敗したユーザーがいます。: %v",
				resObj,
			))
		default:
		}

		c.JSON(http.StatusOK, "")
	})
}
