package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c gcas_account_cost -mockgen ../../../../../bin/mockgen -mock-output mocks/gcas_account_cost_gen.go GCASAccountCost

// GCASAccountCost - GCAS Dashboardから貰ったコスト情報
type GCASAccountCost struct {
	ID        string               `firestore:"-" firestore_key:""` // id
	EventID   string               `firestore:"event_id"`           // event_id
	AccountId string               `firestore:"account_id"`         // アカウントID
	Data      *GCASAccountCostData `firestore:"data"`               // もらったデータ
	Meta
}
