package entities

// GCASAccountCost - GCAS Dashboardから貰ったコスト情報
type GCASAccountCost struct {
	id      string               // id
	eventID string               // event_id
	data    *GCASAccountCostData // もらったデータ
	meta    *Meta                // メタ
}

// NewGCASAccountCostParam - GCAS Dashboardから貰ったコスト情報作成パラメータ
type NewGCASAccountCostParam struct {
	ID      string               // id
	EventID string               // event_id
	Data    *GCASAccountCostData // もらったデータ
	Meta    *Meta                // メタ
}

// NewGCASAccountCost - GCAS Dashboardから貰ったコスト情報作成
func NewGCASAccountCost(param *NewGCASAccountCostParam) *GCASAccountCost {
	return &GCASAccountCost{
		id:      param.ID,
		eventID: param.EventID,
		data:    param.Data,
		meta:    param.Meta,
	}
}

// ID - ID のゲッター
func (e *GCASAccountCost) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *GCASAccountCost) EventID() string {
	return e.eventID
}

// Data - Data のゲッター
func (e *GCASAccountCost) Data() *GCASAccountCostData {
	return e.data
}

// Meta - Meta のゲッター
func (e *GCASAccountCost) Meta() *Meta {
	return e.meta
}
