package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/payment_gen.go Payment

// Payment - 収納情報
type Payment struct {
	ID           string `firestore:"-" firestore_key:""` // id
	EventID      string ``                               // event_id
	BillingID    string ``                               //
	PaymentNumID string ``                               // 収納番号 document ID
	Cost         int    ``                               // この収納番号で請求している金額
	Meta
}
