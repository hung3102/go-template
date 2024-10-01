package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// Billing - 請求
type Billing struct {
	id                   valueobjects.BillingID             // ID
	eventID              valueobjects.EventID               // イベントID
	organization         string                             // 組織名
	csp                  string                             // CSP
	email                string                             // メールアドレス
	address              string                             // 住所
	cost                 int                                // コスト
	orgCSPAccountCostIDs []valueobjects.OrgCSPAccountCostID // アカウントID
	creatorID            valueobjects.CreatorID             // 請求書の発行元情報ID
	billingType          int                                // 支払い種別
	meta                 *Meta                              // メタ
}

// NewBillingParam - 請求作成パラメータ
type NewBillingParam struct {
	ID                   valueobjects.BillingID             // ID
	EventID              valueobjects.EventID               // イベントID
	Organization         string                             // 組織名
	CSP                  string                             // CSP
	Email                string                             // メールアドレス
	Address              string                             // 住所
	Cost                 int                                // コスト
	OrgCSPAccountCostIDs []valueobjects.OrgCSPAccountCostID // アカウントID
	CreatorID            valueobjects.CreatorID             // 請求書の発行元情報ID
	BillingType          int                                // 支払い種別
	Meta                 *Meta                              // メタ
}

// NewBilling - 請求作成
func NewBilling(param *NewBillingParam) *Billing {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewBillingID()
	}
	return &Billing{
		id:                   id,
		eventID:              param.EventID,
		organization:         param.Organization,
		csp:                  param.CSP,
		email:                param.Email,
		address:              param.Address,
		cost:                 param.Cost,
		orgCSPAccountCostIDs: param.OrgCSPAccountCostIDs,
		creatorID:            param.CreatorID,
		billingType:          param.BillingType,
		meta:                 param.Meta,
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

// OrgCSPAccountCostIDs - OrgCSPAccountCostIDs のゲッター
func (e *Billing) OrgCSPAccountCostIDs() []valueobjects.OrgCSPAccountCostID {
	return e.orgCSPAccountCostIDs
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
