package expose

import (
	stripeAPI "github.com/stripe/stripe-go/v74"
	stripeAPISubs "github.com/stripe/stripe-go/v74/subscription"
	"github.com/techstart35/discord-auth-bot/src/server/domain/model"
	"github.com/techstart35/discord-auth-bot/src/server/domain/model/stripe"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"os"
)

// レスポンスです
type Response struct {
	ID             string
	SubscriberID   string
	OperatorRoleID []string
	CustomerID     string
	SubscriptionID string
	Status         string
}

// ステータスを取得します
//
// 取得できない/freeプラン の場合は空の値を返します。
func GetStatus(subscriptionID string) string {
	if subscriptionID == "" {
		return ""
	}

	// Stripe APIキーを設定
	stripeAPI.Key = os.Getenv("STRIPE_SECRET_KEY")

	// サブスクリプション情報を取得
	sub, err := stripeAPISubs.Get(subscriptionID, nil)
	if err != nil {
		return ""
	}

	return string(sub.Status)
}

// サーバーを復元します
func restoreServer(
	id model.ID,
	operatorRoles []string,
	subscriberID string,
	customerID string,
	subscriptionID string,
) (model.Server, error) {
	roles := make([]model.RoleID, 0)
	for _, or := range operatorRoles {
		rID, err := model.NewRoleID(or)
		if err != nil {
			return model.Server{}, errors.NewError("ロールIDを作成できません", err)
		}

		roles = append(roles, rID)
	}

	subscriber, err := model.NewUserID(subscriberID)
	if err != nil {
		return model.Server{}, errors.NewError("支払い者を作成できません", err)
	}

	cusID, err := stripe.NewCustomerID(customerID)
	if err != nil {
		return model.Server{}, errors.NewError("カスタマーIDを作成できません", err)
	}

	subscription, err := stripe.NewSubscriptionID(subscriptionID)
	if err != nil {
		return model.Server{}, errors.NewError("サブスクリプションIDを作成できません", err)
	}

	strp, err := stripe.NewStripe(cusID, subscription)
	if err != nil {
		return model.Server{}, errors.NewError("ストライプを作成できません", err)
	}

	sv := model.RestoreServer(id, roles, subscriber, strp)
	if err != nil {
		return model.Server{}, errors.NewError("サーバー構造体を復元できません", err)
	}

	return sv, nil
}
