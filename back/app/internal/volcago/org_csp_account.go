package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/org_csp_account_gen.go OrgCSPAccount

// OrgCSPAccount - 団体ごとのCSPごとのアカウントごと
type OrgCSPAccount struct {
	ID                   string `firestore:"-" firestore_key:""` // id
	EventID              string ``                               // event_id
	GCASProportionCostID string ``                               // gcas_proportion_cost_id
	GCASAccountCostID    string ``                               // gcas_account_cost_id
	Organization         string ``                               // 団体名
	CSP                  string ``                               // CSP
	AccountID            string ``                               // アカウントID
	Cost                 int    ``                               // 費用按分類
	BillingUnitID        string ``                               // 支払い区分ID
	Meta
}
