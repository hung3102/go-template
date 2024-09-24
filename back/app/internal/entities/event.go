package entities

// Event - イベント
type Event struct {
	id             string // id
	billingMonth   string // 請求月 (例：202408)
	executionCount int    // 何回目の実行か
	meta           *Meta  // メタ
}

// NewEventParam - イベント作成パラメータ
type NewEventParam struct {
	ID             string // id
	BillingMonth   string // 請求月 (例：202408)
	ExecutionCount int    // 何回目の実行か
	Meta           *Meta  // メタ
}

// NewEvent - イベント作成
func NewEvent(param *NewEventParam) *Event {
	return &Event{
		id:             param.ID,
		billingMonth:   param.BillingMonth,
		executionCount: param.ExecutionCount,
		meta:           param.Meta,
	}
}

// ID - ID のゲッター
func (e *Event) ID() string {
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
