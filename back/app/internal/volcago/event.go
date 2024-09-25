package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c event -mockgen ../../../../../bin/mockgen -mock-output mocks/event_gen.go Event

// Event - イベント
type Event struct {
	ID             string `firestore:"-" firestore_key:""` // id
	BillingMonth   string `firestore:"billing_month"`      // 請求月 (例：202408)
	ExecutionCount int    `firestore:"execution_count"`    // 何回目の実行か
	Meta
}
