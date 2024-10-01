package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c email_transaction -mockgen ../../../../../bin/mockgen -mock-output mocks/email_transaction_gen.go EmailTransaction

// EmailTransaction - email_transaction
type EmailTransaction struct {
	ID        string `firestore:"-" firestore_key:""` // ID
	EventID   string `firestore:"event_id"`           // イベントID
	BillingID string `firestore:"billing_id"`         // 請求ID
	Meta
}
