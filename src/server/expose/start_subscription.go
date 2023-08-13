package expose

import (
	"github.com/techstart35/discord-auth-bot/src/server/domain/model"
	"github.com/techstart35/discord-auth-bot/src/server/gateway"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

// TODO: ここを実装

// サブスクリプションを開始した時の処理です
//
// webhookのAPIからコールされます。
func StartSubscription(id, subscriberID, customerID, subscriptionID string) error {
	// idでサーバーを取得
	i, err := model.NewID(id)
	if err != nil {
		return errors.NewError("IDを作成できません", err)
	}

	// IDでDBから情報を取得します
	s, err := gateway.FindByID(i)
	if err != nil {
		return errors.NewError("IDでサーバーを取得できません", err)
	}

	// カスタマーID, subscriptionIDを登録
	sv, err := restoreServer(
		i, s.OperatorRoleID, subscriberID, customerID, subscriptionID,
	)
	if err != nil {
		return errors.NewError("サーバー構造体を復元できません", err)
	}

	if err = gateway.Update(sv); err != nil {
		return errors.NewError("更新できません", err)
	}

	return nil
}
