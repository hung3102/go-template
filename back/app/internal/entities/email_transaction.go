package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// EmailTransaction - email_transaction
type EmailTransaction struct {
	id        valueobjects.EmailTransactionID // id
	eventID   valueobjects.EventID            // event_id
	billingID valueobjects.BillingID          // billing_id
	meta      *Meta                           // メタ
}

// NewEmailTransactionParam - email_transaction作成パラメータ
type NewEmailTransactionParam struct {
	ID        valueobjects.EmailTransactionID // id
	EventID   valueobjects.EventID            // event_id
	BillingID valueobjects.BillingID          // billing_id
	Meta      *Meta                           // メタ
}

// NewEmailTransaction - email_transaction作成
func NewEmailTransaction(param *NewEmailTransactionParam) *EmailTransaction {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewEmailTransactionID()
	}
	return &EmailTransaction{
		id:        id,
		eventID:   param.EventID,
		billingID: param.BillingID,
		meta:      param.Meta,
	}
}

// ID - ID のゲッター
func (e *EmailTransaction) ID() valueobjects.EmailTransactionID {
	return e.id
}

// EventID - EventID のゲッター
func (e *EmailTransaction) EventID() valueobjects.EventID {
	return e.eventID
}

// BillingID - BillingID のゲッター
func (e *EmailTransaction) BillingID() valueobjects.BillingID {
	return e.billingID
}

// Meta - Meta のゲッター
func (e *EmailTransaction) Meta() *Meta {
	return e.meta
}
