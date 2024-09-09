package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=./mocks/service_mock.go -package=mocks

// Provider - 認証サービス
type Provider interface {
	// IssueAuthToken - AuthToken を発行する
	IssueAuthToken(userID, sessionID string, now time.Time) (string, error)
	// VerifyAuthToken - AuthToken を検証する
	VerifyAuthToken(token string) (claims jwt.Claims, err error)
}
