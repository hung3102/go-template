package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c gcas_proportion_cost -mockgen ../../../../../bin/mockgen -mock-output mocks/gcas_proportion_cost_gen.go GCASProportionCost

// GCASProportionCost - GCASから貰った費用按分情報
type GCASProportionCost struct {
	ID        string                  `firestore:"-" firestore_key:""` // ID
	EventID   string                  `firestore:"event_id"`           // イベントID
	AccountID string                  `firestore:"account_id"`         // アカウントID
	Data      *GCASProportionCostData `firestore:"data"`               // GCAS費用按分情報
	Meta
}
