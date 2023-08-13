package webhook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74/webhook"
	"github.com/techstart35/discord-auth-bot/src/server/expose"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"io"
	"net/http"
	"os"
)

// Stripeからコールされるwebhookです
func Webhook(e *gin.Engine) {
	e.POST("/api/webhook", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("リクエストボディの読み取りに失敗しました:", err)
			c.JSON(http.StatusInternalServerError, "内部エラー")
			return
		}

		header := c.GetHeader("Stripe-Signature")
		webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

		event, err := webhook.ConstructEvent(body, header, webhookSecret)
		if err != nil {
			fmt.Println("ここです: 2")
			c.JSON(http.StatusBadRequest, "不正なリクエストです")
			return
		}

		// イベントオブジェクトのdocument
		// https://billing.stripe.com/p/login/test_00geXT3KO9jS2CAaEE
		switch event.Type {
		case "checkout.session.completed":
			fmt.Println("EVENT: completed")
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
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, "サブスクリプションの開始情報を作成できません")
				return
			}
		case "invoice.paid":
			fmt.Println("EVENT: Paied")
			// TODO: 実装
			// 請求期間ごとに、支払いが成功すると送信されます。
			// ステータスをDBに保存します。
		case "customer.subscription.deleted":
			fmt.Println("EVENT: deleted")
			// 顧客のサブスクリプションが終了すると送信されます。
			metadata := event.Data.Object["metadata"].(map[string]string)
			guildID := metadata["guild_id"]

			if err = expose.DeleteSubscription(guildID); err != nil {
				c.JSON(http.StatusInternalServerError, "サブスクリプションの開始情報を作成できません")
				return
			}
		case "invoice.payment_failed":
			fmt.Println("EVENT: failed")
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
