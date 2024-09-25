package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c org_csp_account -mockgen ../../../../../bin/mockgen -mock-output mocks/org_csp_account_gen.go OrgCSPAccount

// OrgCSPAccount - 団体ごとのCSPごとのアカウントごと
type OrgCSPAccount struct {
	ID                   string `firestore:"-" firestore_key:""`      // id
	EventID              string `firestore:"event_id"`                // event_id
	GCASProportionCostID string `firestore:"gcas_proportion_cost_id"` // gcas_proportion_cost_id
	GCASAccountCostID    string `firestore:"gcas_account_cost_id"`    // gcas_account_cost_id
	Organization         string `firestore:"organization"`            // 団体名
	CSP                  string `firestore:"csp"`                     // CSP
	AccountID            string `firestore:"account_id"`              // アカウントID
	Cost                 int    `firestore:"cost"`                    // 費用按分類
	BillingUnitID        string `firestore:"billing_unit_id"`         // 支払い区分ID
	Meta
}
