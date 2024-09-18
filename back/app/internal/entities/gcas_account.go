package entities

type GCASAccount struct {
	id         string // ID
	eventDocID string // event_doc_id
	accountID  string // CSPのアカウントID
}

// NewGCASAccountParam - GCASAccount作成パラメータ
type NewGCASAccountParam struct {
	ID         string // ID
	EventDocID string // event_doc_id
	AccountID  string // CSPのアカウントID
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
