package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c payment_num -mockgen ../../../../../bin/mockgen -mock-output mocks/payment_num_gen.go PaymentNum

// PaymentNum - 収納番号
type PaymentNum struct {
	ID            string `firestore:"-" firestore_key:""` // ID
	EventID       string `firestore:"event_id"`           // イベントID
	BillingID     string `firestore:"billing_id"`         // 請求ID
	BillingUnitID string `firestore:"billing_unit_id"`    // 支払い区分ID
	PaymentNum    string `firestore:"payment_num"`        // 収納番号
	Cost          int    `firestore:"cost"`               // 金額
	Meta
}
