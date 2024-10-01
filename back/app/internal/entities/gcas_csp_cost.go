package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// GCASCSPCost - GCAS Dashboardから貰ったCSPのトータルコスト
type GCASCSPCost struct {
	id        valueobjects.GCASCSPCostID // id
	eventID   valueobjects.EventID       // event_id
	csp       string                     // AWSなど
	totalCost int                        //
	meta      *Meta                      // メタ
}

// NewGCASCSPCostParam - GCAS Dashboardから貰ったCSPのトータルコスト作成パラメータ
type NewGCASCSPCostParam struct {
	ID        valueobjects.GCASCSPCostID // id
	EventID   valueobjects.EventID       // event_id
	CSP       string                     // AWSなど
	TotalCost int                        //
	Meta      *Meta                      // メタ
}

// NewGCASCSPCost - GCAS Dashboardから貰ったCSPのトータルコスト作成
func NewGCASCSPCost(param *NewGCASCSPCostParam) *GCASCSPCost {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewGCASCSPCostID()
	}
	return &GCASCSPCost{
		id:        id,
		eventID:   param.EventID,
		csp:       param.CSP,
		totalCost: param.TotalCost,
		meta:      param.Meta,
	}
}

// ID - ID のゲッター
func (e *GCASCSPCost) ID() valueobjects.GCASCSPCostID {
	return e.id
}

// EventID - EventID のゲッター
func (e *GCASCSPCost) EventID() valueobjects.EventID {
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
