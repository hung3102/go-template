package entities

// Billing - 請求
type Billing struct {
	id                string   // id
	eventID           string   // event_id
	organization      string   //
	csp               string   //
	email             string   //
	address           string   //
	cost              int      //
	orgCSPAccountsIDs []string //
	creatorID         string   //
	billingType       int      //
	meta              *Meta    // メタ
}

// NewBillingParam - 請求作成パラメータ
type NewBillingParam struct {
	ID                string   // id
	EventID           string   // event_id
	Organization      string   //
	CSP               string   //
	Email             string   //
	Address           string   //
	Cost              int      //
	OrgCSPAccountsIDs []string //
	CreatorID         string   //
	BillingType       int      //
	Meta              *Meta    // メタ
}

// NewBilling - 請求作成
func NewBilling(param *NewBillingParam) *Billing {
	return &Billing{
		id:                param.ID,
		eventID:           param.EventID,
		organization:      param.Organization,
		csp:               param.CSP,
		email:             param.Email,
		address:           param.Address,
		cost:              param.Cost,
		orgCSPAccountsIDs: param.OrgCSPAccountsIDs,
		creatorID:         param.CreatorID,
		billingType:       param.BillingType,
		meta:              param.Meta,
	}
}

// ID - ID のゲッター
func (e *Billing) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *Billing) EventID() string {
	return e.eventID
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

// OrgCSPAccountsIDs - OrgCSPAccountsIDs のゲッター
func (e *Billing) OrgCSPAccountsIDs() []string {
	return e.orgCSPAccountsIDs
}

// CreatorID - CreatorID のゲッター
func (e *Billing) CreatorID() string {
	return e.creatorID
}

// BillingType - BillingType のゲッター
func (e *Billing) BillingType() int {
	return e.billingType
}

// Meta - Meta のゲッター
func (e *Billing) Meta() *Meta {
	return e.meta
}
