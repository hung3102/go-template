package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c org_csp_account_cost -mockgen ../../../../../bin/mockgen -mock-output mocks/org_csp_account_cost_gen.go OrgCSPAccountCost

// OrgCSPAccountCost - 団体ごとのCSPごとのアカウントごと
type OrgCSPAccountCost struct {
	ID                   string         `firestore:"-" firestore_key:""`      // ID
	EventID              string         `firestore:"event_id"`                // イベントID
	GCASProportionCostID string         `firestore:"gcas_proportion_cost_id"` // GCAS按分コストID
	GCASAccountCostID    string         `firestore:"gcas_account_cost_id"`    // GCASアカウントコストID
	Organization         string         `firestore:"organization"`            // 団体名
	CSP                  string         `firestore:"csp"`                     // CSP
	AccountID            string         `firestore:"account_id"`              // アカウントID
	Cost                 int            `firestore:"cost"`                    // 金額
	BillingUnitID        string         `firestore:"billing_unit_id"`         // 支払い区分ID
	PaymentAgency        *PaymentAgency `firestore:"payment_agency"`          // 支払い代行者情報
	Meta
}
