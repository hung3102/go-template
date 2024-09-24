package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/payment_num_gen.go PaymentNum

// PaymentNum - 収納番号
type PaymentNum struct {
	ID            string `firestore:"-" firestore_key:""` // id
	EventID       string ``                               // event_id
	BillingID     string ``                               //
	BillingUnitID string ``                               //
	PaymentNum    string ``                               // 収納番号
	Cost          int    ``                               // この収納番号で請求している金額
	Meta
}
