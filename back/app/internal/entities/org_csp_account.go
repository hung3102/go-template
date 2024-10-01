package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// OrgCSPAccount - 団体ごとのCSPごとのアカウントごと
type OrgCSPAccount struct {
	id                   valueobjects.OrgCSPAccountID      // ID
	eventID              valueobjects.EventID              // イベントID
	gcasProportionCostID valueobjects.GCASProportionCostID // GCAS按分コストID
	gcasAccountCostID    valueobjects.GCASAccountCostID    // GCASアカウントコストID
	organization         string                            // 団体名
	csp                  string                            // CSP
	accountID            string                            // アカウントID
	cost                 int                               // 費用按分類
	billingUnitID        string                            // 支払い区分ID
	meta                 *Meta                             // メタ
}

// NewOrgCSPAccountParam - 団体ごとのCSPごとのアカウントごと作成パラメータ
type NewOrgCSPAccountParam struct {
	ID                   valueobjects.OrgCSPAccountID      // ID
	EventID              valueobjects.EventID              // イベントID
	GCASProportionCostID valueobjects.GCASProportionCostID // GCAS按分コストID
	GCASAccountCostID    valueobjects.GCASAccountCostID    // GCASアカウントコストID
	Organization         string                            // 団体名
	CSP                  string                            // CSP
	AccountID            string                            // アカウントID
	Cost                 int                               // 費用按分類
	BillingUnitID        string                            // 支払い区分ID
	Meta                 *Meta                             // メタ
}

// NewOrgCSPAccount - 団体ごとのCSPごとのアカウントごと作成
func NewOrgCSPAccount(param *NewOrgCSPAccountParam) *OrgCSPAccount {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewOrgCSPAccountID()
	}
	return &OrgCSPAccount{
		id:                   id,
		eventID:              param.EventID,
		gcasProportionCostID: param.GCASProportionCostID,
		gcasAccountCostID:    param.GCASAccountCostID,
		organization:         param.Organization,
		csp:                  param.CSP,
		accountID:            param.AccountID,
		cost:                 param.Cost,
		billingUnitID:        param.BillingUnitID,
		meta:                 param.Meta,
	}
}

// ID - ID のゲッター
func (e *OrgCSPAccount) ID() valueobjects.OrgCSPAccountID {
	return e.id
}

// EventID - EventID のゲッター
func (e *OrgCSPAccount) EventID() valueobjects.EventID {
	return e.eventID
}

// GCASProportionCostID - GCASProportionCostID のゲッター
func (e *OrgCSPAccount) GCASProportionCostID() valueobjects.GCASProportionCostID {
	return e.gcasProportionCostID
}

// GCASAccountCostID - GCASAccountCostID のゲッター
func (e *OrgCSPAccount) GCASAccountCostID() valueobjects.GCASAccountCostID {
	return e.gcasAccountCostID
}

// Organization - Organization のゲッター
func (e *OrgCSPAccount) Organization() string {
	return e.organization
}

// CSP - CSP のゲッター
func (e *OrgCSPAccount) CSP() string {
	return e.csp
}

// AccountID - AccountID のゲッター
func (e *OrgCSPAccount) AccountID() string {
	return e.accountID
}

// Cost - Cost のゲッター
func (e *OrgCSPAccount) Cost() int {
	return e.cost
}

// BillingUnitID - BillingUnitID のゲッター
func (e *OrgCSPAccount) BillingUnitID() string {
	return e.billingUnitID
}

// Meta - Meta のゲッター
func (e *OrgCSPAccount) Meta() *Meta {
	return e.meta
}
