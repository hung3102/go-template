package volcago

import "time"

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/user_session_gen.go UserSession

// UserSession - ユーザーセッション
type UserSession struct {
	ID        string    `firestore:"-" firestore_key:""` // ID
	UserID    string    ``                               // ユーザーID
	ExpiresAt time.Time ``                               // 有効期限
	Meta
}
