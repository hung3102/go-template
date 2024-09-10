package event

import "time"

// Event - イベント
type Event struct {
	ID           string    // イベントID
	BillingMonth time.Time // 請求月
	Status       string    // ステータス
	CreatedAt    time.Time // 作成日時
}
