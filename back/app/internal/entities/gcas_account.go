package entities

// GCASAccount - GCASアカウント情報
type GCASAccount struct {
	id         string // ID
	eventDocID string // event_doc_id
	accountID  string // CSPのアカウントID
	meta       *Meta  // メタ
}

// NewGCASAccountParam - GCASAccount作成パラメータ
type NewGCASAccountParam struct {
	ID         string // ID
	EventDocID string // event_doc_id
	AccountID  string // CSPのアカウントID
	Meta       *Meta  // Meta
}

// GCASAccount - GCASAccount作成
func NewGCASAccount(param *NewGCASAccountParam) *GCASAccount {
	return &GCASAccount{
		id:         param.ID,
		eventDocID: param.EventDocID,
		accountID:  param.AccountID,
	}
}

// ID - ID のゲッター
func (e *GCASAccount) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *GCASAccount) EventDocID() string {
	return e.eventDocID
}

// AccountID - AccountID のゲッター
func (e *GCASAccount) AccountID() string {
	return e.accountID
}

// Meta - meta のゲッター
func (e *GCASAccount) Meta() *Meta {
	return e.meta
}
