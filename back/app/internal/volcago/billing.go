package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c billing -mockgen ../../../../../bin/mockgen -mock-output mocks/billing_gen.go Billing

// Billing - 請求
type Billing struct {
	ID                string   `firestore:"-" firestore_key:""`   // id
	EventID           string   `firestore:"event_id"`             // event_id
	Organization      string   `firestore:"organization"`         //
	CSP               string   `firestore:"csp"`                  //
	Email             string   `firestore:"email"`                //
	Address           string   `firestore:"address"`              //
	Cost              int      `firestore:"cost"`                 //
	OrgCSPAccountsIDs []string `firestore:"org_csp_accounts_ids"` //
	CreatorID         string   `firestore:"creator_id"`           //
	BillingType       int      `firestore:"billing_type"`         //
	Meta
}
