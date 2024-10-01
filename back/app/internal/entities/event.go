package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// Event - イベント
type Event struct {
	id             valueobjects.EventID // ID
	billingMonth   string               // 請求月
	executionCount int                  // 実行回数
	meta           *Meta                // メタ
}

// NewEventParam - イベント作成パラメータ
type NewEventParam struct {
	ID             valueobjects.EventID // ID
	BillingMonth   string               // 請求月
	ExecutionCount int                  // 実行回数
	Meta           *Meta                // メタ
}

// NewEvent - イベント作成
func NewEvent(param *NewEventParam) *Event {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewEventID()
	}
	return &Event{
		id:             id,
		billingMonth:   param.BillingMonth,
		executionCount: param.ExecutionCount,
		meta:           param.Meta,
	}
}

// ID - ID のゲッター
func (e *Event) ID() valueobjects.EventID {
	return e.id
}

// BillingMonth - BillingMonth のゲッター
func (e *Event) BillingMonth() string {
	return e.billingMonth
}

// ExecutionCount - ExecutionCount のゲッター
func (e *Event) ExecutionCount() int {
	return e.executionCount
}

// Meta - Meta のゲッター
func (e *Event) Meta() *Meta {
	return e.meta
}
