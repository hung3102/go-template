package entities

const (
	EventStatusStart                  = 1 // 1. 開始
	EventStatusInvoiceCreationChecked = 2 // 2.請求作成可能確認済
	EventStatusProportionStart        = 3 // 3. 按分処理開始
	EventStatusProportionCompleted    = 4 // 4. 按分処理済
)

// EventStatus - イベントステータス
type EventStatus struct {
	id         string // id
	eventDocID string // event_doc_id
	status     int    // ステータス
	meta       *Meta  // メタ
}

// NewEventStatusParam - イベントステータス作成パラメータ
type NewEventStatusParam struct {
	ID         string // id
	EventDocID string // event_doc_id
	Status     int    // ステータス
	Meta       *Meta  // Meta
}

// NewEventStatus - イベントステータス作成
func NewEventStatus(param *NewEventStatusParam) *EventStatus {
	return &EventStatus{
		id:         param.ID,
		eventDocID: param.EventDocID,
		status:     param.Status,
		meta:       param.Meta,
	}
}

// ID - ID のゲッター
func (e *EventStatus) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *EventStatus) EventDocID() string {
	return e.eventDocID
}

// Status - Status のゲッター
func (e *EventStatus) Status() int {
	return e.status
}

// Meta - meta のゲッター
func (e *EventStatus) Meta() *Meta {
	return e.meta
}
