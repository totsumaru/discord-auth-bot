package model

import "github.com/techstart35/discord-auth-bot/src/shared/errors"

const LimitOperatorRoleAmount = 3

// サーバーです
type Server struct {
	id             ID       // サーバーID
	subscriberID   UserID   // 支払い者のユーザーID
	operatorRoleID []RoleID // 操作できるロールID
}

// サーバーを作成します
func NewServer(id ID, subsID UserID, operator []RoleID) (Server, error) {
	s := Server{}
	s.id = id
	s.subscriberID = subsID
	s.operatorRoleID = operator

	return s, nil
}

// 支払い者を変更します
func (s *Server) UpdateSubscriberID(newSubsID UserID) error {
	s.subscriberID = newSubsID

	if err := s.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// 操作ロールを変更します
func (s *Server) UpdateOperatorRoleID(newOperator []RoleID) error {
	s.operatorRoleID = newOperator

	if err := s.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// サーバーIDを取得します
func (s *Server) ID() ID {
	return s.id
}

// 支払い者のユーザーIDを取得します
func (s *Server) SubscriberID() UserID {
	return s.subscriberID
}

// オペレーターロールIDを取得します
func (s *Server) OperatorRoleID() []RoleID {
	return s.operatorRoleID
}

// 検証します
func (s *Server) validate() error {
	if len(s.operatorRoleID) > LimitOperatorRoleAmount {
		return errors.NewError("オペレーターロールの上限を超えています")
	}

	return nil
}
