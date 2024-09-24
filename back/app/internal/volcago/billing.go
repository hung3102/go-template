package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/billing_gen.go Billing

// Billing - 請求
type Billing struct {
	ID                string   `firestore:"-" firestore_key:""` // id
	EventID           string   ``                               // event_id
	Organization      string   ``                               //
	CSP               string   ``                               //
	Email             string   ``                               //
	Address           string   ``                               //
	Cost              int      ``                               //
	OrgCSPAccountsIDs []string ``                               //
	CreatorID         string   ``                               //
	BillingType       int      ``                               //
	Meta
}
