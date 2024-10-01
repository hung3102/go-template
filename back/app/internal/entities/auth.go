package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// Auth - REPSアクセストークン
type Auth struct {
	id           valueobjects.AuthID // ID
	eventId      string              // イベントID
	accessToken  string              // アクセストークン
	refreshToken string              // リフレッシュトークン
	meta         *Meta               // メタ
}

// NewAuthParam - REPSアクセストークン作成パラメータ
type NewAuthParam struct {
	ID           valueobjects.AuthID // ID
	EventId      string              // イベントID
	AccessToken  string              // アクセストークン
	RefreshToken string              // リフレッシュトークン
	Meta         *Meta               // メタ
}

// NewAuth - REPSアクセストークン作成
func NewAuth(param *NewAuthParam) *Auth {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewAuthID()
	}
	return &Auth{
		id:           id,
		eventId:      param.EventId,
		accessToken:  param.AccessToken,
		refreshToken: param.RefreshToken,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *Auth) ID() valueobjects.AuthID {
	return e.id
}

// EventId - EventId のゲッター
func (e *Auth) EventId() string {
	return e.eventId
}

// AccessToken - AccessToken のゲッター
func (e *Auth) AccessToken() string {
	return e.accessToken
}

// RefreshToken - RefreshToken のゲッター
func (e *Auth) RefreshToken() string {
	return e.refreshToken
}

// Meta - Meta のゲッター
func (e *Auth) Meta() *Meta {
	return e.meta
}
