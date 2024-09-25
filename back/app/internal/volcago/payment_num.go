package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c payment_num -mockgen ../../../../../bin/mockgen -mock-output mocks/payment_num_gen.go PaymentNum

// PaymentNum - 収納番号
type PaymentNum struct {
	ID            string `firestore:"-" firestore_key:""` // id
	EventID       string `firestore:"event_id"`           // event_id
	BillingID     string `firestore:"billing_id"`         //
	BillingUnitID string `firestore:"billing_unit_id"`    //
	PaymentNum    string `firestore:"payment_num"`        // 収納番号
	Cost          int    `firestore:"cost"`               // この収納番号で請求している金額
	Meta
}
