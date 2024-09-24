package entities

// Payment - 収納情報
type Payment struct {
	id           string // id
	eventID      string // event_id
	billingID    string //
	paymentNumID string // 収納番号 document ID
	cost         int    // この収納番号で請求している金額
	meta         *Meta  // メタ
}

// NewPaymentParam - 収納情報作成パラメータ
type NewPaymentParam struct {
	ID           string // id
	EventID      string // event_id
	BillingID    string //
	PaymentNumID string // 収納番号 document ID
	Cost         int    // この収納番号で請求している金額
	Meta         *Meta  // メタ
}

// NewPayment - 収納情報作成
func NewPayment(param *NewPaymentParam) *Payment {
	return &Payment{
		id:           param.ID,
		eventID:      param.EventID,
		billingID:    param.BillingID,
		paymentNumID: param.PaymentNumID,
		cost:         param.Cost,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *Payment) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *Payment) EventID() string {
	return e.eventID
}

// BillingID - BillingID のゲッター
func (e *Payment) BillingID() string {
	return e.billingID
}

// PaymentNumID - PaymentNumID のゲッター
func (e *Payment) PaymentNumID() string {
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
