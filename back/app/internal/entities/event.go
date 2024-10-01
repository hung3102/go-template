package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// Event - イベント
type Event struct {
	id             valueobjects.EventID // id
	billingMonth   string               // 請求月 (例：202408)
	executionCount int                  // 何回目の実行か
	meta           *Meta                // メタ
}

// NewEventParam - イベント作成パラメータ
type NewEventParam struct {
	ID             valueobjects.EventID // id
	BillingMonth   string               // 請求月 (例：202408)
	ExecutionCount int                  // 何回目の実行か
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
