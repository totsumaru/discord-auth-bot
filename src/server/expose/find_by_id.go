package expose

import (
	"github.com/techstart35/discord-auth-bot/src/server/domain/model"
	"github.com/techstart35/discord-auth-bot/src/server/gateway"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

// IDでサーバー情報を取得します
func FindByID(id string) (Response, error) {
	res := Response{}

	i, err := model.NewID(id)
	if err != nil {
		return res, errors.NewError("IDを作成できません", err)
	}

	s, err := gateway.FindByID(i)
	if err != nil {
		return res, errors.NewError("IDでサーバーを取得できません", err)
	}

	res.ID = s.ID
	res.SubscriberID = s.SubscriberID
	res.OperatorRoleID = s.OperatorRoleID
	res.CustomerID = s.CustomerID
	res.SubscriptionID = s.SubscriptionID
	res.Status = GetStatus(s.SubscriptionID)

	return res, nil
}
