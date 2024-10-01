package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c reps_auth -mockgen ../../../../../bin/mockgen -mock-output mocks/reps_auth_gen.go REPSAuth

// REPSAuth - REPSアクセストークン
type REPSAuth struct {
	ID           string `firestore:"-" firestore_key:""` // ID
	EventID      string `firestore:"event_id"`           // イベントID
	AccessToken  string `firestore:"access_token"`       // アクセストークン
	RefreshToken string `firestore:"refresh_token"`      // リフレッシュトークン
	Meta
}
