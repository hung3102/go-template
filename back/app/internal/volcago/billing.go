package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c billing -mockgen ../../../../../bin/mockgen -mock-output mocks/billing_gen.go Billing

// Billing - 請求
type Billing struct {
	ID                   string   `firestore:"-" firestore_key:""`        // ID
	EventID              string   `firestore:"event_id"`                  // イベントID
	Organization         string   `firestore:"organization"`              // 組織名
	CSP                  string   `firestore:"csp"`                       // CSP
	Email                string   `firestore:"email"`                     // メールアドレス
	Address              string   `firestore:"address"`                   // 住所
	Cost                 int      `firestore:"cost"`                      // コスト
	OrgCSPAccountCostIDs []string `firestore:"org_csp_accounts_cost_ids"` // アカウントID
	CreatorID            string   `firestore:"creator_id"`                // 請求書の発行元情報ID
	BillingType          int      `firestore:"billing_type"`              // 支払い種別
	Meta
}
