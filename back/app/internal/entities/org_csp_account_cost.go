package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// OrgCSPAccountCost - 団体ごとのCSPごとのアカウントごと
type OrgCSPAccountCost struct {
	id                   valueobjects.OrgCSPAccountCostID  // ID
	eventID              valueobjects.EventID              // イベントID
	gcasProportionCostID valueobjects.GCASProportionCostID // GCAS按分コストID
	gcasAccountCostID    valueobjects.GCASAccountCostID    // GCASアカウントコストID
	organization         string                            // 団体名
	csp                  string                            // CSP
	accountID            string                            // アカウントID
	cost                 int                               // 金額
	billingUnitID        string                            // 支払い区分ID
	paymentAgency        *PaymentAgency                    // 支払い代行者情報
	meta                 *Meta                             // メタ
}

// NewOrgCSPAccountCostParam - 団体ごとのCSPごとのアカウントごと作成パラメータ
type NewOrgCSPAccountCostParam struct {
	ID                   valueobjects.OrgCSPAccountCostID  // ID
	EventID              valueobjects.EventID              // イベントID
	GCASProportionCostID valueobjects.GCASProportionCostID // GCAS按分コストID
	GCASAccountCostID    valueobjects.GCASAccountCostID    // GCASアカウントコストID
	Organization         string                            // 団体名
	CSP                  string                            // CSP
	AccountID            string                            // アカウントID
	Cost                 int                               // 金額
	BillingUnitID        string                            // 支払い区分ID
	PaymentAgency        *PaymentAgency                    // 支払い代行者情報
	Meta                 *Meta                             // メタ
}

// NewOrgCSPAccountCost - 団体ごとのCSPごとのアカウントごと作成
func NewOrgCSPAccountCost(param *NewOrgCSPAccountCostParam) *OrgCSPAccountCost {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewOrgCSPAccountCostID()
	}
	return &OrgCSPAccountCost{
		id:                   id,
		eventID:              param.EventID,
		gcasProportionCostID: param.GCASProportionCostID,
		gcasAccountCostID:    param.GCASAccountCostID,
		organization:         param.Organization,
		csp:                  param.CSP,
		accountID:            param.AccountID,
		cost:                 param.Cost,
		billingUnitID:        param.BillingUnitID,
		paymentAgency:        param.PaymentAgency,
		meta:                 param.Meta,
	}
}

// ID - ID のゲッター
func (e *OrgCSPAccountCost) ID() valueobjects.OrgCSPAccountCostID {
	return e.id
}

// EventID - EventID のゲッター
func (e *OrgCSPAccountCost) EventID() valueobjects.EventID {
	return e.eventID
}

// GCASProportionCostID - GCASProportionCostID のゲッター
func (e *OrgCSPAccountCost) GCASProportionCostID() valueobjects.GCASProportionCostID {
	return e.gcasProportionCostID
}

// GCASAccountCostID - GCASAccountCostID のゲッター
func (e *OrgCSPAccountCost) GCASAccountCostID() valueobjects.GCASAccountCostID {
	return e.gcasAccountCostID
}

// Organization - Organization のゲッター
func (e *OrgCSPAccountCost) Organization() string {
	return e.organization
}

// CSP - CSP のゲッター
func (e *OrgCSPAccountCost) CSP() string {
	return e.csp
}

// AccountID - AccountID のゲッター
func (e *OrgCSPAccountCost) AccountID() string {
	return e.accountID
}

// Cost - Cost のゲッター
func (e *OrgCSPAccountCost) Cost() int {
	return e.cost
}

// BillingUnitID - BillingUnitID のゲッター
func (e *OrgCSPAccountCost) BillingUnitID() string {
	return e.billingUnitID
}

// PaymentAgency - PaymentAgency のゲッター
func (e *OrgCSPAccountCost) PaymentAgency() *PaymentAgency {
	return e.paymentAgency
}

// Meta - Meta のゲッター
func (e *OrgCSPAccountCost) Meta() *Meta {
	return e.meta
}
