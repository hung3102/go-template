package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c auth -mockgen ../../../../../bin/mockgen -mock-output mocks/auth_gen.go Auth

// Auth - REPSアクセストークン
type Auth struct {
	ID           string `firestore:"-" firestore_key:""` // ID
	EventId      string `firestore:"event_id"`           // イベントID
	AccessToken  string `firestore:"access_token"`       // アクセストークン
	RefreshToken string `firestore:"refresh_token"`      // リフレッシュトークン
	Meta
}
