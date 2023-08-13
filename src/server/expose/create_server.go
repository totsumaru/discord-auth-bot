package expose

import (
	"github.com/techstart35/discord-auth-bot/src/server/domain/model"
	"github.com/techstart35/discord-auth-bot/src/server/gateway"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

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
