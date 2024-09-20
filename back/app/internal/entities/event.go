package entities

import (
	"time"
)

// Event - イベント
type Event struct {
	id           string    // id
	billingMonth time.Time // 請求月 (YYYYMM)
	meta         *Meta     // メタ
}

// NewEventParam - イベント作成パラメータ
type NewEventParam struct {
	ID           string    // id
	BillingMonth time.Time // 請求月 (YYYYMM)
	Meta         *Meta     // Meta
}

// NewEvent - イベント作成
func NewEvent(param *NewEventParam) *Event {
	return &Event{
		id:           param.ID,
		billingMonth: param.BillingMonth,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *Event) ID() string {
	return e.id
}

// BillingMonth - BillingMonth のゲッター
func (e *Event) BillingMonth() time.Time {
	return e.billingMonth
}

// Meta - meta のゲッター
func (e *Event) Meta() *Meta {
	return e.meta
}
