package shared

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	errors "github.com/techstart35/discord-auth-bot/src/shared"
	seeker "github.com/techstart35/discord-auth-bot/src/shared/map"
	"os"
	"strings"
)

const HeaderAuthorization = "Authorization"

// TODO: 有効期限の確認を行う

// Header(Bearer xxx)からdiscordIDを取得します
// ヘッダーの取得例) authHeader := c.GetHeader(shared.HeaderAuthorization)
func GetDiscordIDFromAuthHeader(authHeader string) (string, error) {
	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	// トークンを解析
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("メソッドが期待した値と異なります: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", errors.NewError("認証できません")
	}

	// トークンが有効なら、ユーザーはログインしていると判断できる
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := seeker.Str(claims, []string{"user_metadata", "provider_id"})
		return id, nil
	} else {
		return "", errors.NewError("トークンが無効です")
	}
}
