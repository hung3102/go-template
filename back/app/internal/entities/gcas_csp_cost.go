package entities

type GCASCSPCost struct {
	id         string // ID
	eventDocID string // event_doc_id
	csp        string // CSP名
	totalCost  int    // コスト合計
	meta       *Meta  // メタ
}

// NewGCASCSPCostParam - GCASCSPCost作成パラメータ
type NewGCASCSPCostParam struct {
	ID         string // ID
	EventDocID string // event_doc_id
	CSP        string // CSP名
	TotalCost  int    // コスト合計
	Meta       *Meta  // Meta
}

// GCASCSPCost - GCASCSPCost作成
func NewGCASCSPCost(param *NewGCASCSPCostParam) *GCASCSPCost {
	return &GCASCSPCost{
		id:         param.ID,
		eventDocID: param.EventDocID,
		csp:        param.CSP,
		totalCost:  param.TotalCost,
	}
}

// ID - ID のゲッター
func (e *GCASCSPCost) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *GCASCSPCost) EventDocID() string {
	return e.eventDocID
}

// CSP - CSP のゲッター
func (e *GCASCSPCost) CSP() string {
	return e.csp
}

// TotalCost - TotalCost のゲッター
func (e *GCASCSPCost) TotalCost() int {
	return e.totalCost
}

// Meta - meta のゲッター
func (e *GCASCSPCost) Meta() *Meta {
	return e.meta
}
