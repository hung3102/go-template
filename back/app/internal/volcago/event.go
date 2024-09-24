package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/event_gen.go Event

// Event - イベント
type Event struct {
	ID             string `firestore:"-" firestore_key:""` // id
	BillingMonth   string ``                               // 請求月 (例：202408)
	ExecutionCount int    ``                               // 何回目の実行か
	Meta
}
