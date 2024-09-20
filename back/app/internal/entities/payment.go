package entities

// Payment - 収納情報
type Payment struct {
	id              string // id
	eventDocID      string // event_doc_id
	billingDocID    string //
	paymentNumDocID string // 収納番号 document ID
	cost            int    // この収納番号で請求している金額
	meta            *Meta  // メタ
}

// NewPaymentParam - 収納情報作成パラメータ
type NewPaymentParam struct {
	ID              string // id
	EventDocID      string // event_doc_id
	BillingDocID    string //
	PaymentNumDocID string // 収納番号 document ID
	Cost            int    // この収納番号で請求している金額
	Meta            *Meta  // メタ
}

// NewPayment - 収納情報作成
func NewPayment(param *NewPaymentParam) *Payment {
	return &Payment{
		id:              param.ID,
		eventDocID:      param.EventDocID,
		billingDocID:    param.BillingDocID,
		paymentNumDocID: param.PaymentNumDocID,
		cost:            param.Cost,
		meta:            param.Meta,
	}
}

// ID - ID のゲッター
func (e *Payment) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *Payment) EventDocID() string {
	return e.eventDocID
}

// BillingDocID - BillingDocID のゲッター
func (e *Payment) BillingDocID() string {
	return e.billingDocID
}

// PaymentNumDocID - PaymentNumDocID のゲッター
func (e *Payment) PaymentNumDocID() string {
	return e.paymentNumDocID
}

// Cost - Cost のゲッター
func (e *Payment) Cost() int {
	return e.cost
}

// Meta - Meta のゲッター
func (e *Payment) Meta() *Meta {
	return e.meta
}
