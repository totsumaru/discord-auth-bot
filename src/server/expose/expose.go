package expose

import (
	"github.com/techstart35/discord-auth-bot/src/server/domain/model"
	"github.com/techstart35/discord-auth-bot/src/server/gateway"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

// レスポンスです
type Response struct {
	ID             string
	SubscriberID   string
	OperatorRoleID []string
}

// サーバーにbotが追加された時、DBにサーバーを新規作成します
func CreateServer(id string) error {
	i, err := model.NewID(id)
	if err != nil {
		return errors.NewError("IDを作成できません", err)
	}

	s, err := model.NewServer(i)
	if err != nil {
		return errors.NewError("サーバーを作成できません", err)
	}

	if err = gateway.Create(s.ID()); err != nil {
		return errors.NewError("サーバーを保存できません", err)
	}

	return nil
}

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

	return res, nil
}

// TODO: 実装/支払いを実行します

// オペレーターロールを変更します
func UpdateOperatorRoleID(id string, roles []string) error {
	i, err := model.NewID(id)
	if err != nil {
		return errors.NewError("IDを作成できません", err)
	}

	s, err := gateway.FindByID(i)
	if err != nil {
		return errors.NewError("IDでサーバーを取得できません", err)
	}

	subsID, err := model.NewUserID(s.SubscriberID)
	if err != nil {
		return errors.NewError("支払い者を作成できません", err)
	}

	operatorRoles := make([]model.RoleID, 0)
	for _, v := range roles {
		rID, err := model.NewRoleID(v)
		if err != nil {
			return errors.NewError("ロールIDを作成できません", err)
		}

		operatorRoles = append(operatorRoles, rID)
	}

	if err = gateway.Update(i, subsID, operatorRoles); err != nil {
		return errors.NewError("更新できません", err)
	}

	return nil
}
