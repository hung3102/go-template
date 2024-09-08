package models

import "time"

//go:generate ../../../../../bin/volcago -p repositories -o ../repositories -mockgen ../../../../../bin/mockgen -mock-output mocks/event_gen.go Event

// Event - イベント
type Event struct {
	ID           string    `firestore:"-" firestore_key:"auto"` // ID
	BillingMonth time.Time ``                                   // 請求月 (YYYYMM)
	Status       string    ``                                   // ステータス
	Meta
}
