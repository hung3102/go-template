package entities

// GCASProportionCostData - GCASからもらった費用按分情報詳細
type GCASProportionCostData struct {
	// TODO
}

// NewGCASProportionCostDataParam - GCASからもらった費用按分情報詳細パラメータ
type NewGCASProportionCostDataParam struct {
}

// NewGCASProportionCostData - GCASからもらった費用按分情報詳細作成
func NewGCASProportionCostData(param *NewGCASProportionCostDataParam) *GCASProportionCostData {
	return &GCASProportionCostData{}
}
