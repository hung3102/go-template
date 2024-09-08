package entities

import "time"

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

// CreatedAt - getter of CreatedAt
func (e *Event) CreatedAt() time.Time {
	return e.meta.createdAt
}

// UpdatedAt - getter of UpdatedAt
func (e *Event) UpdatedAt() time.Time {
	return e.meta.updatedAt
}

// DeletedAt - getter of DeletedAt
func (e *Event) DeletedAt() *time.Time {
	return e.meta.deletedAt
}
