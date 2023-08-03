package model

import (
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"regexp"
)

// ロールIDです
type RoleID struct {
	value string
}

// ロールIDを作成します
func NewRoleID(v string) (RoleID, error) {
	id := RoleID{}
	id.value = v

	if err := id.validate(); err != nil {
		return id, errors.NewError("検証に失敗しました", err)
	}

	return id, nil
}

// ロールIDを取得します
func (i RoleID) Value() string {
	return i.value
}

// 検証をします
func (i RoleID) validate() error {
	match, err := regexp.MatchString(`^[0-9]+$`, i.value)
	if err != nil {
		return errors.NewError("正規表現の検証に失敗しました", err)
	}

	if !match {
		return errors.NewError("数値以外の値が含まれています", err)
	}

	return nil
}
