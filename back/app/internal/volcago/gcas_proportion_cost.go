package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/gcas_proportion_cost_gen.go GCASProportionCost

// GCASProportionCost - GCASから貰った費用按分情報
type GCASProportionCost struct {
	ID      string                  `firestore:"-" firestore_key:""` // id
	EventID string                  ``                               // event_id
	Data    *GCASProportionCostData ``                               // もらったデータ
	Meta
}
