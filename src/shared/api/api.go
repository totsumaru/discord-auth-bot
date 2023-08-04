package api

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"github.com/techstart35/discord-auth-bot/src/shared/map/seeker"
	"os"
	"strings"
)

const HeaderAuthorization = "Authorization"

// TODO: 有効期限の確認を行う

type Res struct {
	DiscordID string
}

// Header(Bearer xxx)からdiscordIDを取得します
// ヘッダーの取得例) authHeader := c.GetHeader(shared.HeaderAuthorization)
func GetAuthHeader(authHeader string) (Res, error) {
	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	// トークンを解析
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("メソッドが期待した値と異なります: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return Res{}, errors.NewError("認証できません", err)
	}

	// トークンが有効なら、ユーザーはログインしていると判断できる
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return Res{
			DiscordID: seeker.Str(claims, []string{"user_metadata", "provider_id"}),
		}, nil
	} else {
		return Res{}, errors.NewError("トークンが無効です")
	}
}
