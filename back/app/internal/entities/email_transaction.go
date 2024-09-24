package entities

// EmailTransaction - email_transaction
type EmailTransaction struct {
	id        string // id
	eventID   string // event_id
	billingID string // billing_id
	meta      *Meta  // メタ
}

// NewEmailTransactionParam - email_transaction作成パラメータ
type NewEmailTransactionParam struct {
	ID        string // id
	EventID   string // event_id
	BillingID string // billing_id
	Meta      *Meta  // メタ
}

// NewEmailTransaction - email_transaction作成
func NewEmailTransaction(param *NewEmailTransactionParam) *EmailTransaction {
	return &EmailTransaction{
		id:        param.ID,
		eventID:   param.EventID,
		billingID: param.BillingID,
		meta:      param.Meta,
	}
}

// ID - ID のゲッター
func (e *EmailTransaction) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *EmailTransaction) EventID() string {
	return e.eventID
}

// BillingID - BillingID のゲッター
func (e *EmailTransaction) BillingID() string {
	return e.billingID
}

// Meta - Meta のゲッター
func (e *EmailTransaction) Meta() *Meta {
	return e.meta
}
