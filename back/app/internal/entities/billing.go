package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// Billing - 請求
type Billing struct {
	id                valueobjects.BillingID         // id
	eventID           valueobjects.EventID           // event_id
	organization      string                         //
	csp               string                         //
	email             string                         //
	address           string                         //
	cost              int                            //
	orgCSPAccountsIDs []valueobjects.OrgCSPAccountID //
	creatorID         valueobjects.CreatorID         //
	billingType       int                            //
	meta              *Meta                          // メタ
}

// NewBillingParam - 請求作成パラメータ
type NewBillingParam struct {
	ID                valueobjects.BillingID         // id
	EventID           valueobjects.EventID           // event_id
	Organization      string                         //
	CSP               string                         //
	Email             string                         //
	Address           string                         //
	Cost              int                            //
	OrgCSPAccountsIDs []valueobjects.OrgCSPAccountID //
	CreatorID         valueobjects.CreatorID         //
	BillingType       int                            //
	Meta              *Meta                          // メタ
}

// NewBilling - 請求作成
func NewBilling(param *NewBillingParam) *Billing {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewBillingID()
	}
	return &Billing{
		id:                id,
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
func (e *Billing) ID() valueobjects.BillingID {
	return e.id
}

// EventID - EventID のゲッター
func (e *Billing) EventID() valueobjects.EventID {
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
func (e *Billing) OrgCSPAccountsIDs() []valueobjects.OrgCSPAccountID {
	return e.orgCSPAccountsIDs
}

// CreatorID - CreatorID のゲッター
func (e *Billing) CreatorID() valueobjects.CreatorID {
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
