package model

import (
	"github.com/techstart35/discord-auth-bot/src/server/domain/model/stripe"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

const LimitOperatorRoleAmount = 3

// サーバーです
type Server struct {
	id             ID            // サーバーID
	operatorRoleID []RoleID      // 操作できるロールID
	subscriberID   UserID        // 支払い者のユーザーID
	stripe         stripe.Stripe // Stripeの情報
}

// サーバーを作成します
func NewServer(id ID) (Server, error) {
	s := Server{}
	s.id = id

	return s, nil
}

// サーバーを復元します
func RestoreServer(
	id ID,
	operatorRoleID []RoleID,
	subscriberID UserID,
	stripe stripe.Stripe,
) Server {
	s := Server{}
	s.id = id
	s.operatorRoleID = operatorRoleID
	s.subscriberID = subscriberID
	s.stripe = stripe

	return s
}

// 操作ロールを変更します
func (s *Server) UpdateOperatorRoleID(newOperator []RoleID) error {
	s.operatorRoleID = newOperator

	if err := s.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// 支払い者を変更します
func (s *Server) UpdateSubscriberID(newSubsID UserID) error {
	s.subscriberID = newSubsID

	if err := s.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// Stripeの情報を変更します
func (s *Server) UpdateStripe(stripe stripe.Stripe) error {
	s.stripe = stripe

	if err := s.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// サーバーIDを取得します
func (s *Server) ID() ID {
	return s.id
}

// オペレーターロールIDを取得します
func (s *Server) OperatorRoleID() []RoleID {
	return s.operatorRoleID
}

// 支払い者のユーザーIDを取得します
func (s *Server) SubscriberID() UserID {
	return s.subscriberID
}

// ストライプの情報を取得します
func (s *Server) Stripe() stripe.Stripe {
	return s.stripe
}

// オペレーターロールIDをstringのsliceで取得します
func (s *Server) OperatorRoleIDByStr() []string {
	res := make([]string, 0)
	for _, v := range s.operatorRoleID {
		res = append(res, v.value)
	}

	return res
}

// 検証します
func (s *Server) validate() error {
	if len(s.operatorRoleID) > LimitOperatorRoleAmount {
		return errors.NewError("オペレーターロールの上限を超えています")
	}

	return nil
}
