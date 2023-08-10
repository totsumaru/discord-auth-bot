package res

import (
	"github.com/techstart35/discord-auth-bot/src/api/_utils/permission"
)

// =======================
// レスポンスの統一規格です
// =======================

// サーバーのレスポンスです
type Server struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

// ロールのレスポンスです
type Role struct {
	ID         string                `json:"id"`
	Name       string                `json:"name"`
	Color      int                   `json:"color"`
	Comment    string                `json:"comment"`
	Permission permission.Permission `json:"permission"`
}

// チャンネルのレスポンスです
type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ユーザーのレスポンスです
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}
