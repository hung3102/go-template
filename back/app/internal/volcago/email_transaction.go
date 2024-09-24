package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/email_transaction_gen.go EmailTransaction

// EmailTransaction - email_transaction
type EmailTransaction struct {
	ID        string `firestore:"-" firestore_key:""` // id
	EventID   string ``                               // event_id
	BillingID string ``                               // billing_id
	Meta
}
