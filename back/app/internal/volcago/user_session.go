package volcago

import "time"

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c user_session -mockgen ../../../../../bin/mockgen -mock-output mocks/user_session_gen.go UserSession

// UserSession - ユーザーセッション
type UserSession struct {
	ID        string    `firestore:"-" firestore_key:""` // ID
	UserID    string    `firestore:"user_id"`            // イベントID
	ExpiresAt time.Time `firestore:"expires_at"`         // 有効期限
	Meta
}
