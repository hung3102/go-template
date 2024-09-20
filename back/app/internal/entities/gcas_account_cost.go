package entities

// GCASAccountCost - GCAS Dashboardから貰ったコスト情報
type GCASAccountCost struct {
	id         string               // id
	eventDocID string               // event_doc_id
	data       *GCASAccountCostData // もらったデータ
	meta       *Meta                // メタ
}

// NewGCASAccountCostParam - GCAS Dashboardから貰ったコスト情報作成パラメータ
type NewGCASAccountCostParam struct {
	ID         string               // id
	EventDocID string               // event_doc_id
	Data       *GCASAccountCostData // もらったデータ
	Meta       *Meta                // メタ
}

// NewGCASAccountCost - GCAS Dashboardから貰ったコスト情報作成
func NewGCASAccountCost(param *NewGCASAccountCostParam) *GCASAccountCost {
	return &GCASAccountCost{
		id:         param.ID,
		eventDocID: param.EventDocID,
		data:       param.Data,
		meta:       param.Meta,
	}
}

// ID - ID のゲッター
func (e *GCASAccountCost) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *GCASAccountCost) EventDocID() string {
	return e.eventDocID
}

// Data - Data のゲッター
func (e *GCASAccountCost) Data() *GCASAccountCostData {
	return e.data
}

// Meta - Meta のゲッター
func (e *GCASAccountCost) Meta() *Meta {
	return e.meta
}
