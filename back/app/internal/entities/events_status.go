package entities

// EventsStatus - イベントステータス
type EventsStatus struct {
	id      string // {event_id}_{status}
	eventID string // event_id
	status  int    // ステータス
	meta    *Meta  // メタ
}

// NewEventsStatusParam - イベントステータス作成パラメータ
type NewEventsStatusParam struct {
	ID      string // {event_id}_{status}
	EventID string // event_id
	Status  int    // ステータス
	Meta    *Meta  // メタ
}

// NewEventsStatus - イベントステータス作成
func NewEventsStatus(param *NewEventsStatusParam) *EventsStatus {
	return &EventsStatus{
		id:      param.ID,
		eventID: param.EventID,
		status:  param.Status,
		meta:    param.Meta,
	}
}

// ID - ID のゲッター
func (e *EventsStatus) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *EventsStatus) EventID() string {
	return e.eventID
}

// Status - Status のゲッター
func (e *EventsStatus) Status() int {
	return e.status
}

// Meta - Meta のゲッター
func (e *EventsStatus) Meta() *Meta {
	return e.meta
}
