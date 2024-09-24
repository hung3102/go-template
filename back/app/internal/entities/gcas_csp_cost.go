package entities

// GCASCSPCost - GCAS Dashboardから貰ったCSPのトータルコスト
type GCASCSPCost struct {
	id        string // id
	eventID   string // event_id
	csp       string // AWSなど
	totalCost int    //
	meta      *Meta  // メタ
}

// NewGCASCSPCostParam - GCAS Dashboardから貰ったCSPのトータルコスト作成パラメータ
type NewGCASCSPCostParam struct {
	ID        string // id
	EventID   string // event_id
	CSP       string // AWSなど
	TotalCost int    //
	Meta      *Meta  // メタ
}

// NewGCASCSPCost - GCAS Dashboardから貰ったCSPのトータルコスト作成
func NewGCASCSPCost(param *NewGCASCSPCostParam) *GCASCSPCost {
	return &GCASCSPCost{
		id:        param.ID,
		eventID:   param.EventID,
		csp:       param.CSP,
		totalCost: param.TotalCost,
		meta:      param.Meta,
	}
}

// ID - ID のゲッター
func (e *GCASCSPCost) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *GCASCSPCost) EventID() string {
	return e.eventID
}

// CSP - CSP のゲッター
func (e *GCASCSPCost) CSP() string {
	return e.csp
}

// TotalCost - TotalCost のゲッター
func (e *GCASCSPCost) TotalCost() int {
	return e.totalCost
}

// Meta - Meta のゲッター
func (e *GCASCSPCost) Meta() *Meta {
	return e.meta
}
