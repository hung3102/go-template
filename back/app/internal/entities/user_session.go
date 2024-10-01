package entities

import (
	"time"

	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
)

// UserSession - ユーザーセッション
type UserSession struct {
	id        valueobjects.UserSessionID // ID
	userID    string                     // ユーザーID
	expiresAt time.Time                  // 有効期限
	meta      *Meta                      // メタ
}

// NewUserSessionParam - ユーザーセッション作成パラメータ
type NewUserSessionParam struct {
	ID        valueobjects.UserSessionID // ID
	UserID    string                     // ユーザーID
	ExpiresAt time.Time                  // 有効期限
	Meta      *Meta                      // メタ
}

// NewUserSession - ユーザーセッション作成
func NewUserSession(param *NewUserSessionParam) *UserSession {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewUserSessionID()
	}
	return &UserSession{
		id:        id,
		userID:    param.UserID,
		expiresAt: param.ExpiresAt,
		meta:      param.Meta,
	}
}

// ID - ID のゲッター
func (e *UserSession) ID() valueobjects.UserSessionID {
	return e.id
}

// UserID - UserID のゲッター
func (e *UserSession) UserID() string {
	return e.userID
}

// ExpiresAt - ExpiresAt のゲッター
func (e *UserSession) ExpiresAt() time.Time {
	return e.expiresAt
}

// Meta - Meta のゲッター
func (e *UserSession) Meta() *Meta {
	return e.meta
}
