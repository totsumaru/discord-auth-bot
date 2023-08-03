package model

import (
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"regexp"
)

// ユーザーIDです
type UserID struct {
	value string
}

// ユーザーIDを作成します
func NewUserID(v string) (UserID, error) {
	id := UserID{}
	id.value = v

	if err := id.validate(); err != nil {
		return id, errors.NewError("検証に失敗しました", err)
	}

	return id, nil
}

// ユーザーIDを取得します
func (i UserID) Value() string {
	return i.value
}

// 検証をします
func (i UserID) validate() error {
	if i.value == "" {
		return nil
	}

	match, err := regexp.MatchString(`^[0-9]+$`, i.value)
	if err != nil {
		return errors.NewError("正規表現の検証に失敗しました", err)
	}

	if !match {
		return errors.NewError("数値以外の値が含まれています", err)
	}

	return nil
}
