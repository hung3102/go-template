package entities

// Billing - 請求
type Billing struct {
	id                   string   // id
	eventDocID           string   // event_doc_id
	orgCSPDocID          string   //
	organization         string   //
	csp                  string   //
	email                string   //
	address              string   //
	cost                 int      //
	orgCSPAccountsDocIDs []string //
	creatorDocID         string   //
	billingType          int      //
	meta                 *Meta    // メタ
}

// NewBillingParam - 請求作成パラメータ
type NewBillingParam struct {
	ID                   string   // id
	EventDocID           string   // event_doc_id
	OrgCSPDocID          string   //
	Organization         string   //
	CSP                  string   //
	Email                string   //
	Address              string   //
	Cost                 int      //
	OrgCSPAccountsDocIDs []string //
	CreatorDocID         string   //
	BillingType          int      //
	Meta                 *Meta    // メタ
}

// NewBilling - 請求作成
func NewBilling(param *NewBillingParam) *Billing {
	return &Billing{
		id:                   param.ID,
		eventDocID:           param.EventDocID,
		orgCSPDocID:          param.OrgCSPDocID,
		organization:         param.Organization,
		csp:                  param.CSP,
		email:                param.Email,
		address:              param.Address,
		cost:                 param.Cost,
		orgCSPAccountsDocIDs: param.OrgCSPAccountsDocIDs,
		creatorDocID:         param.CreatorDocID,
		billingType:          param.BillingType,
		meta:                 param.Meta,
	}
}

// ID - ID のゲッター
func (e *Billing) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *Billing) EventDocID() string {
	return e.eventDocID
}

// OrgCSPDocID - OrgCSPDocID のゲッター
func (e *Billing) OrgCSPDocID() string {
	return e.orgCSPDocID
}

// Organization - Organization のゲッター
func (e *Billing) Organization() string {
	return e.organization
}

// CSP - CSP のゲッター
func (e *Billing) CSP() string {
	return e.csp
}

// Email - Email のゲッター
func (e *Billing) Email() string {
	return e.email
}

// Address - Address のゲッター
func (e *Billing) Address() string {
	return e.address
}

// Cost - Cost のゲッター
func (e *Billing) Cost() int {
	return e.cost
}

// OrgCSPAccountsDocIDs - OrgCSPAccountsDocIDs のゲッター
func (e *Billing) OrgCSPAccountsDocIDs() []string {
	return e.orgCSPAccountsDocIDs
}

// CreatorDocID - CreatorDocID のゲッター
func (e *Billing) CreatorDocID() string {
	return e.creatorDocID
}

// BillingType - BillingType のゲッター
func (e *Billing) BillingType() int {
	return e.billingType
}

// Meta - Meta のゲッター
func (e *Billing) Meta() *Meta {
	return e.meta
}
