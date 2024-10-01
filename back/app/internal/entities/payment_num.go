package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// PaymentNum - 収納番号
type PaymentNum struct {
	id            valueobjects.PaymentNumID // id
	eventID       valueobjects.EventID      // event_id
	billingID     valueobjects.BillingID    //
	billingUnitID string                    // 支払い区分ID
	paymentNum    string                    // 収納番号
	cost          int                       // この収納番号で請求している金額
	meta          *Meta                     // メタ
}

// NewPaymentNumParam - 収納番号作成パラメータ
type NewPaymentNumParam struct {
	ID            valueobjects.PaymentNumID // id
	EventID       valueobjects.EventID      // event_id
	BillingID     valueobjects.BillingID    //
	BillingUnitID string                    // 支払い区分ID
	PaymentNum    string                    // 収納番号
	Cost          int                       // この収納番号で請求している金額
	Meta          *Meta                     // メタ
}

// NewPaymentNum - 収納番号作成
func NewPaymentNum(param *NewPaymentNumParam) *PaymentNum {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewPaymentNumID()
	}
	return &PaymentNum{
		id:            id,
		eventID:       param.EventID,
		billingID:     param.BillingID,
		billingUnitID: param.BillingUnitID,
		paymentNum:    param.PaymentNum,
		cost:          param.Cost,
		meta:          param.Meta,
	}
}

// ID - ID のゲッター
func (e *PaymentNum) ID() valueobjects.PaymentNumID {
	return e.id
}

// EventID - EventID のゲッター
func (e *PaymentNum) EventID() valueobjects.EventID {
	return e.eventID
}

// BillingID - BillingID のゲッター
func (e *PaymentNum) BillingID() valueobjects.BillingID {
	return e.billingID
}

// BillingUnitID - BillingUnitID のゲッター
func (e *PaymentNum) BillingUnitID() string {
	return e.billingUnitID
}

// PaymentNum - PaymentNum のゲッター
func (e *PaymentNum) PaymentNum() string {
	return e.paymentNum
}

// Cost - Cost のゲッター
func (e *PaymentNum) Cost() int {
	return e.cost
}

// Meta - Meta のゲッター
func (e *PaymentNum) Meta() *Meta {
	return e.meta
}
