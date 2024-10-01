package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// REPSAuth - REPSアクセストークン
type REPSAuth struct {
	id           valueobjects.REPSAuthID // ID
	eventID      string                  // イベントID
	accessToken  string                  // アクセストークン
	refreshToken string                  // リフレッシュトークン
	meta         *Meta                   // メタ
}

// NewREPSAuthParam - REPSアクセストークン作成パラメータ
type NewREPSAuthParam struct {
	ID           valueobjects.REPSAuthID // ID
	EventID      string                  // イベントID
	AccessToken  string                  // アクセストークン
	RefreshToken string                  // リフレッシュトークン
	Meta         *Meta                   // メタ
}

// NewREPSAuth - REPSアクセストークン作成
func NewREPSAuth(param *NewREPSAuthParam) *REPSAuth {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewREPSAuthID()
	}
	return &REPSAuth{
		id:           id,
		eventID:      param.EventID,
		accessToken:  param.AccessToken,
		refreshToken: param.RefreshToken,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *REPSAuth) ID() valueobjects.REPSAuthID {
	return e.id
}

// EventID - EventID のゲッター
func (e *REPSAuth) EventID() string {
	return e.eventID
}

// AccessToken - AccessToken のゲッター
func (e *REPSAuth) AccessToken() string {
	return e.accessToken
}

// RefreshToken - RefreshToken のゲッター
func (e *REPSAuth) RefreshToken() string {
	return e.refreshToken
}

// Meta - Meta のゲッター
func (e *REPSAuth) Meta() *Meta {
	return e.meta
}
