package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c gcas_csp_cost -mockgen ../../../../../bin/mockgen -mock-output mocks/gcas_csp_cost_gen.go GCASCSPCost

// GCASCSPCost - GCAS Dashboardから貰ったCSPのトータルコスト
type GCASCSPCost struct {
	ID        string `firestore:"-" firestore_key:""` // id
	EventID   string `firestore:"event_id"`           // event_id
	CSP       string `firestore:"csp"`                // AWSなど
	TotalCost int    `firestore:"total_cost"`         //
	Meta
}
