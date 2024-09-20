package entities

// OrgCSP - 団体ごとのCSPごとの情報
type OrgCSP struct {
	id           string // id
	eventDocID   string // event_doc_id
	email        string // email
	organization string // 団体名
	csp          string // CSP
	meta         *Meta  // メタ
}

// NewOrgCSPParam - 団体ごとのCSPごとの情報作成パラメータ
type NewOrgCSPParam struct {
	ID           string // id
	EventDocID   string // event_doc_id
	Email        string // email
	Organization string // 団体名
	CSP          string // CSP
	Meta         *Meta  // メタ
}

// NewOrgCSP - 団体ごとのCSPごとの情報作成
func NewOrgCSP(param *NewOrgCSPParam) *OrgCSP {
	return &OrgCSP{
		id:           param.ID,
		eventDocID:   param.EventDocID,
		email:        param.Email,
		organization: param.Organization,
		csp:          param.CSP,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *OrgCSP) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *OrgCSP) EventDocID() string {
	return e.eventDocID
}

// Email - Email のゲッター
func (e *OrgCSP) Email() string {
	return e.email
}

// Organization - Organization のゲッター
func (e *OrgCSP) Organization() string {
	return e.organization
}

// CSP - CSP のゲッター
func (e *OrgCSP) CSP() string {
	return e.csp
}

// Meta - Meta のゲッター
func (e *OrgCSP) Meta() *Meta {
	return e.meta
}
