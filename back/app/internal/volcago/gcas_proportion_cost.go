package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c gcas_proportion_cost -mockgen ../../../../../bin/mockgen -mock-output mocks/gcas_proportion_cost_gen.go GCASProportionCost

// GCASProportionCost - GCASから貰った費用按分情報
type GCASProportionCost struct {
	ID        string                  `firestore:"-" firestore_key:""` //
	EventID   string                  `firestore:"event_id"`           //
	AccountID string                  `firestore:"account_id"`         //
	Data      *GCASProportionCostData `firestore:"data"`               // GCAS費用按分情報APIから取得
	Meta
}
