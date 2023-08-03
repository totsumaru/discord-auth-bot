package model

import (
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"regexp"
)

// サーバーIDです
type ID struct {
	value string
}

// サーバーIDを作成します
func NewID(v string) (ID, error) {
	id := ID{}
	id.value = v

	if err := id.validate(); err != nil {
		return id, errors.NewError("検証に失敗しました", err)
	}

	return id, nil
}

// サーバーIDを取得します
func (i ID) Value() string {
	return i.value
}

// 検証をします
func (i ID) validate() error {
	match, err := regexp.MatchString(`^[0-9]+$`, i.value)
	if err != nil {
		return errors.NewError("正規表現の検証に失敗しました", err)
	}

	if !match {
		return errors.NewError("数値以外の値が含まれています", err)
	}

	return nil
}
