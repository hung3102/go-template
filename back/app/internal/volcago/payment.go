package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c payment -mockgen ../../../../../bin/mockgen -mock-output mocks/payment_gen.go Payment

// Payment - 収納情報
type Payment struct {
	ID           string `firestore:"-" firestore_key:""` // ID
	EventID      string `firestore:"event_id"`           // イベントID
	BillingID    string `firestore:"billing_id"`         // 請求ID
	PaymentNumID string `firestore:"payment_num_id"`     // 収納番号
	Cost         int    `firestore:"cost"`               // 金額
	Meta
}
