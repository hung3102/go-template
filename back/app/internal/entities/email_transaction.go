package entities

// EmailTransaction - email_transaction
type EmailTransaction struct {
	id          string // id
	eventDocID  string // event_doc_id
	orgCspDocID string // org_csp_doc_id
	meta        *Meta  // メタ
}

// NewEmailTransactionParam - email_transaction作成パラメータ
type NewEmailTransactionParam struct {
	ID          string // id
	EventDocID  string // event_doc_id
	OrgCspDocID string // org_csp_doc_id
	Meta        *Meta  // メタ
}

// NewEmailTransaction - email_transaction作成
func NewEmailTransaction(param *NewEmailTransactionParam) *EmailTransaction {
	return &EmailTransaction{
		id:          param.ID,
		eventDocID:  param.EventDocID,
		orgCspDocID: param.OrgCspDocID,
		meta:        param.Meta,
	}
}

// ID - ID のゲッター
func (e *EmailTransaction) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *EmailTransaction) EventDocID() string {
	return e.eventDocID
}

// OrgCspDocID - OrgCspDocID のゲッター
func (e *EmailTransaction) OrgCspDocID() string {
	return e.orgCspDocID
}

// Meta - Meta のゲッター
func (e *EmailTransaction) Meta() *Meta {
	return e.meta
}
