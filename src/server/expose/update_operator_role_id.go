package expose

import (
	"github.com/techstart35/discord-auth-bot/src/server/domain/model"
	"github.com/techstart35/discord-auth-bot/src/server/gateway"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

// オペレーターロールを変更します
func UpdateOperatorRoleID(id string, roles []string) error {
	i, err := model.NewID(id)
	if err != nil {
		return errors.NewError("IDを作成できません", err)
	}

	// IDでDBから情報を取得します
	s, err := gateway.FindByID(i)
	if err != nil {
		return errors.NewError("IDでサーバーを取得できません", err)
	}

	sv, err := restoreServer(
		i, roles, s.SubscriberID, s.CustomerID, s.SubscriptionID,
	)
	if err != nil {
		return errors.NewError("サーバー構造体を復元できません", err)
	}

	if err = gateway.Update(sv); err != nil {
		return errors.NewError("更新できません", err)
	}

	return nil
}
