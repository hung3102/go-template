package entities

// GCASProportionCost - GCASから貰った費用按分情報
type GCASProportionCost struct {
	id         string                  // id
	eventDocID string                  // event_doc_id
	data       *GCASProportionCostData // もらったデータ
	meta       *Meta                   // メタ
}

// NewGCASProportionCostParam - GCASから貰った費用按分情報作成パラメータ
type NewGCASProportionCostParam struct {
	ID         string                  // id
	EventDocID string                  // event_doc_id
	Data       *GCASProportionCostData // もらったデータ
	Meta       *Meta                   // メタ
}

// NewGCASProportionCost - GCASから貰った費用按分情報作成
func NewGCASProportionCost(param *NewGCASProportionCostParam) *GCASProportionCost {
	return &GCASProportionCost{
		id:         param.ID,
		eventDocID: param.EventDocID,
		data:       param.Data,
		meta:       param.Meta,
	}
}

// ID - ID のゲッター
func (e *GCASProportionCost) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *GCASProportionCost) EventDocID() string {
	return e.eventDocID
}

// Data - Data のゲッター
func (e *GCASProportionCost) Data() *GCASProportionCostData {
	return e.data
}

// Meta - Meta のゲッター
func (e *GCASProportionCost) Meta() *Meta {
	return e.meta
}
