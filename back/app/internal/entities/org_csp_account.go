package entities

// ORGCSPAccount - 団体ごとのCSPごとのアカウントごと
type ORGCSPAccount struct {
	id                      string // id
	eventDocID              string // event_doc_id
	gcasProportionCostDocID string // gcas_proportion_cost_doc_id
	gcasAccountCostDocID    string // gcas_account_cost_doc_id
	organization            string // 団体名
	csp                     string // CSP
	accountID               string // アカウントID
	cost                    int    // 費用按分類
	billingUnitID           string // 支払い区分ID
	meta                    *Meta  // メタ
}

// NewORGCSPAccountParam - 団体ごとのCSPごとのアカウントごと作成パラメータ
type NewORGCSPAccountParam struct {
	ID                      string // id
	EventDocID              string // event_doc_id
	GCASProportionCostDocID string // gcas_proportion_cost_doc_id
	GCASAccountCostDocID    string // gcas_account_cost_doc_id
	Organization            string // 団体名
	CSP                     string // CSP
	AccountID               string // アカウントID
	Cost                    int    // 費用按分類
	BillingUnitID           string // 支払い区分ID
	Meta                    *Meta  // メタ
}

// NewORGCSPAccount - 団体ごとのCSPごとのアカウントごと作成
func NewORGCSPAccount(param *NewORGCSPAccountParam) *ORGCSPAccount {
	return &ORGCSPAccount{
		id:                      param.ID,
		eventDocID:              param.EventDocID,
		gcasProportionCostDocID: param.GCASProportionCostDocID,
		gcasAccountCostDocID:    param.GCASAccountCostDocID,
		organization:            param.Organization,
		csp:                     param.CSP,
		accountID:               param.AccountID,
		cost:                    param.Cost,
		billingUnitID:           param.BillingUnitID,
		meta:                    param.Meta,
	}
}

// ID - ID のゲッター
func (e *ORGCSPAccount) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *ORGCSPAccount) EventDocID() string {
	return e.eventDocID
}

// GCASProportionCostDocID - GCASProportionCostDocID のゲッター
func (e *ORGCSPAccount) GCASProportionCostDocID() string {
	return e.gcasProportionCostDocID
}

// GCASAccountCostDocID - GCASAccountCostDocID のゲッター
func (e *ORGCSPAccount) GCASAccountCostDocID() string {
	return e.gcasAccountCostDocID
}

// Organization - Organization のゲッター
func (e *ORGCSPAccount) Organization() string {
	return e.organization
}

// CSP - CSP のゲッター
func (e *ORGCSPAccount) CSP() string {
	return e.csp
}

// AccountID - AccountID のゲッター
func (e *ORGCSPAccount) AccountID() string {
	return e.accountID
}

// Cost - Cost のゲッター
func (e *ORGCSPAccount) Cost() int {
	return e.cost
}

// BillingUnitID - BillingUnitID のゲッター
func (e *ORGCSPAccount) BillingUnitID() string {
	return e.billingUnitID
}

// Meta - Meta のゲッター
func (e *ORGCSPAccount) Meta() *Meta {
	return e.meta
}
