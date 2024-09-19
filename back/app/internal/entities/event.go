package entities

import (
	"strconv"
	"time"
)

const (
	// EventStatusInvoiceCreationPossible - 請求書作成可能
	EventStatusInvoiceCreationPossible = 1
	// EventStatusStored - 収納済
	EventStatusStored = 2
)

// Event - イベント
type Event struct {
	id           string    // id
	billingMonth time.Time // 請求月 (YYYYMM)
	status       string    // ステータス
	meta         *Meta     // メタ
}

// NewEventParam - イベント作成パラメータ
type NewEventParam struct {
	ID           string    // id
	BillingMonth time.Time // 請求月 (YYYYMM)
	Status       string    // ステータス
	Meta         *Meta     // Meta
}

// NewEvent - イベント作成
func NewEvent(param *NewEventParam) *Event {
	return &Event{
		id:           param.ID,
		billingMonth: param.BillingMonth,
		status:       param.Status,
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

// Status - Status のゲッター
func (e *Event) Status() string {
	return e.status
}

// Meta - meta のゲッター
func (e *Event) Meta() *Meta {
	return e.meta
}

// IsInvoiceCreateionPossible - 請求作成可能か判定する。
func (e *Event) IsInvoiceCreateionPossible() bool {
	return e.status == strconv.Itoa(EventStatusInvoiceCreationPossible)
}
