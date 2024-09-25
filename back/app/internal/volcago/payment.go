package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c payment -mockgen ../../../../../bin/mockgen -mock-output mocks/payment_gen.go Payment

// Payment - 収納情報
type Payment struct {
	ID           string `firestore:"-" firestore_key:""` // id
	EventID      string `firestore:"event_id"`           // event_id
	BillingID    string `firestore:"billing_id"`         //
	PaymentNumID string `firestore:"payment_num_id"`     // 収納番号 document ID
	Cost         int    `firestore:"cost"`               // この収納番号で請求している金額
	Meta
}
