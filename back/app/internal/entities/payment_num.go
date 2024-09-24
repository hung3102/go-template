package entities

// PaymentNum - 収納番号
type PaymentNum struct {
	id            string // id
	eventID       string // event_id
	billingID     string //
	billingUnitID string //
	paymentNum    string // 収納番号
	cost          int    // この収納番号で請求している金額
	meta          *Meta  // メタ
}

// NewPaymentNumParam - 収納番号作成パラメータ
type NewPaymentNumParam struct {
	ID            string // id
	EventID       string // event_id
	BillingID     string //
	BillingUnitID string //
	PaymentNum    string // 収納番号
	Cost          int    // この収納番号で請求している金額
	Meta          *Meta  // メタ
}

// NewPaymentNum - 収納番号作成
func NewPaymentNum(param *NewPaymentNumParam) *PaymentNum {
	return &PaymentNum{
		id:            param.ID,
		eventID:       param.EventID,
		billingID:     param.BillingID,
		billingUnitID: param.BillingUnitID,
		paymentNum:    param.PaymentNum,
		cost:          param.Cost,
		meta:          param.Meta,
	}
}

// ID - ID のゲッター
func (e *PaymentNum) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *PaymentNum) EventID() string {
	return e.eventID
}

// BillingID - BillingID のゲッター
func (e *PaymentNum) BillingID() string {
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
