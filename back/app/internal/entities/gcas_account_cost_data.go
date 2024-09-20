package entities

// GCASAccountCostData - GCAS Dashboardから貰ったコスト情報詳細
type GCASAccountCostData struct {
	// TODO
}

// NewGCASAccountCostDataParam - GCAS Dashboardから貰ったコスト情報詳細作成パラメータ
type NewGCASAccountCostDataParam struct {
	ID         string               // id
	EventDocID string               // event_doc_id
	Data       *GCASAccountCostData // もらったデータ
	Meta       *Meta                // Meta

}

// NewGCASAccountCostData - GCAS Dashboardから貰ったコスト情報詳細作成
func NewGCASAccountCostData(param *NewGCASAccountCostDataParam) *GCASAccountCostData {
	return &GCASAccountCostData{}
}
