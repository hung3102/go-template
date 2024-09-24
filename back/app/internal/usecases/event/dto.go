package event

import "time"

// Event - イベント
type Event struct {
	ID           string    // イベントID
	BillingMonth string    // 請求月
	CreatedAt    time.Time // 作成日時
}
