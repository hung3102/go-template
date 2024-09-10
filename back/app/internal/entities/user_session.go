package entities

import "time"

// UserSession - ユーザセッション
type UserSession struct {
	id        string    // id
	userID    string    // ユーザID
	expiresAt time.Time // 有効期限
	meta      *Meta     // メタ
}

// NewUserSessionParam - NewUserSession のパラメータ
type NewUserSessionParam struct {
	ID        string    // id
	UserID    string    // ユーザID
	ExpiresAt time.Time // 有効期限
	Meta      *Meta     // Meta
}

// NewUserSession - UserSession のコンストラクタ
func NewUserSession(param *NewUserSessionParam) *UserSession {
	return &UserSession{
		id:        param.ID,
		userID:    param.UserID,
		expiresAt: param.ExpiresAt,
		meta:      param.Meta,
	}
}

// ID - id のゲッター
func (u *UserSession) ID() string {
	return u.id
}

// UserID - userID のゲッター
func (u *UserSession) UserID() string {
	return u.userID
}

// ExpiresAt - expiresAt のゲッター
func (u *UserSession) ExpiresAt() time.Time {
	return u.expiresAt
}

// Meta - meta のゲッター
func (u *UserSession) Meta() *Meta {
	return u.meta
}
