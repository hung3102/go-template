package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/gcas_csp_cost_gen.go GCASCSPCost

// GCASCSPCost - GCAS Dashboardから貰ったCSPのトータルコスト
type GCASCSPCost struct {
	ID        string `firestore:"-" firestore_key:""` // id
	EventID   string ``                               // event_id
	CSP       string ``                               // AWSなど
	TotalCost int    ``                               //
	Meta
}
