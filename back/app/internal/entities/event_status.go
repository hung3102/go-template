package entities

const (
	EventStatusStart                  = 1 // 1. 開始
	EventStatusInvoiceCreationChecked = 2 // 2.請求作成可能確認済
	EventStatusProportionStart        = 3 // 3. 按分処理開始
	EventStatusProportionCompleted    = 4 // 4. 按分処理済
)

// EventStatus - イベントステータス
type EventStatus struct {
	id      string // {event_id}_{status}
	eventID string // event_id
	status  int    // ステータス
	meta    *Meta  // メタ
}

// NewEventStatusParam - イベントステータス作成パラメータ
type NewEventStatusParam struct {
	ID      string // {event_id}_{status}
	EventID string // event_id
	Status  int    // ステータス
	Meta    *Meta  // メタ
}

// NewEventStatus - イベントステータス作成
func NewEventStatus(param *NewEventStatusParam) *EventStatus {
	return &EventStatus{
		id:      param.ID,
		eventID: param.EventID,
		status:  param.Status,
		meta:    param.Meta,
	}
}

// ID - ID のゲッター
func (e *EventStatus) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *EventStatus) EventID() string {
	return e.eventID
}

// Status - Status のゲッター
func (e *EventStatus) Status() int {
	return e.status
}

// Meta - Meta のゲッター
func (e *EventStatus) Meta() *Meta {
	return e.meta
}
