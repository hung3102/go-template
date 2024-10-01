package entities

import (
	"fmt"

	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
)

const (
	EventStatusStart                  = 1 // 1. 開始
	EventStatusInvoiceCreationChecked = 2 // 2.請求作成可能確認済
	EventStatusProportionStart        = 3 // 3. 按分処理開始
	EventStatusProportionCompleted    = 4 // 4. 按分処理済
)

// EventStatus - イベントステータス
type EventStatus struct {
	id      string               // {event_id}_{status}
	eventID valueobjects.EventID // イベントID
	status  int                  // ステータス
	meta    *Meta                // メタ
}

// NewEventStatusParam - イベントステータス作成パラメータ
type NewEventStatusParam struct {
	EventID valueobjects.EventID // イベントID
	Status  int                  // ステータス
	Meta    *Meta                // メタ
}

// NewEventStatus - イベントステータス作成
func NewEventStatus(param *NewEventStatusParam) *EventStatus {
	return &EventStatus{
		id:      ToEventStatusID(param),
		eventID: param.EventID,
		status:  param.Status,
		meta:    param.Meta,
	}
}

// ToEventStatusID - event_statusのIDを取得する
func ToEventStatusID(param *NewEventStatusParam) string {
	return fmt.Sprintf("%s_%d", param.EventID, param.Status)
}

// ID - ID のゲッター
func (e *EventStatus) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *EventStatus) EventID() valueobjects.EventID {
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
