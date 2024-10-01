package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// Payment - 収納情報
type Payment struct {
	id           valueobjects.PaymentID    // id
	eventID      valueobjects.EventID      // event_id
	billingID    valueobjects.BillingID    //
	paymentNumID valueobjects.PaymentNumID // 収納番号 document ID
	cost         int                       // この収納番号で請求している金額
	meta         *Meta                     // メタ
}

// NewPaymentParam - 収納情報作成パラメータ
type NewPaymentParam struct {
	ID           valueobjects.PaymentID    // id
	EventID      valueobjects.EventID      // event_id
	BillingID    valueobjects.BillingID    //
	PaymentNumID valueobjects.PaymentNumID // 収納番号 document ID
	Cost         int                       // この収納番号で請求している金額
	Meta         *Meta                     // メタ
}

// NewPayment - 収納情報作成
func NewPayment(param *NewPaymentParam) *Payment {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewPaymentID()
	}
	return &Payment{
		id:           id,
		eventID:      param.EventID,
		billingID:    param.BillingID,
		paymentNumID: param.PaymentNumID,
		cost:         param.Cost,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *Payment) ID() valueobjects.PaymentID {
	return e.id
}

// EventID - EventID のゲッター
func (e *Payment) EventID() valueobjects.EventID {
	return e.eventID
}

// BillingID - BillingID のゲッター
func (e *Payment) BillingID() valueobjects.BillingID {
	return e.billingID
}

// PaymentNumID - PaymentNumID のゲッター
func (e *Payment) PaymentNumID() valueobjects.PaymentNumID {
	return e.paymentNumID
}

// Cost - Cost のゲッター
func (e *Payment) Cost() int {
	return e.cost
}

// Meta - Meta のゲッター
func (e *Payment) Meta() *Meta {
	return e.meta
}
