package jwtimpl

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims - クレーム
type Claims struct {
	UserID    string `json:"userID"`    // ユーザーID
	SessionID string `json:"sessionID"` // セッションID
	jwt.RegisteredClaims
}
