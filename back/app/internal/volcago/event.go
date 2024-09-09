package volcago

import "time"

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/event_gen.go Event

// Event - イベント
type Event struct {
	ID           string    `firestore:"-" firestore_key:""` // ID
	BillingMonth time.Time ``                               // 請求月 (YYYYMM)
	Status       string    ``                               // ステータス
	Meta
}
