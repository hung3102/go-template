package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/gcas_account_cost_gen.go GCASAccountCost

// GCASAccountCost - GCAS Dashboardから貰ったコスト情報
type GCASAccountCost struct {
	ID      string               `firestore:"-" firestore_key:""` // id
	EventID string               ``                               // event_id
	Data    *GCASAccountCostData ``                               // もらったデータ
	Meta
}
