package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// GCASProportionCost - GCASから貰った費用按分情報
type GCASProportionCost struct {
	id        valueobjects.GCASProportionCostID // ID
	eventID   valueobjects.EventID              // イベントID
	accountID string                            // アカウントID
	data      *GCASProportionCostData           // GCAS費用按分情報
	meta      *Meta                             // メタ
}

// NewGCASProportionCostParam - GCASから貰った費用按分情報作成パラメータ
type NewGCASProportionCostParam struct {
	ID        valueobjects.GCASProportionCostID // ID
	EventID   valueobjects.EventID              // イベントID
	AccountID string                            // アカウントID
	Data      *GCASProportionCostData           // GCAS費用按分情報
	Meta      *Meta                             // メタ
}

// NewGCASProportionCost - GCASから貰った費用按分情報作成
func NewGCASProportionCost(param *NewGCASProportionCostParam) *GCASProportionCost {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewGCASProportionCostID()
	}
	return &GCASProportionCost{
		id:        id,
		eventID:   param.EventID,
		accountID: param.AccountID,
		data:      param.Data,
		meta:      param.Meta,
	}
}

// ID - ID のゲッター
func (e *GCASProportionCost) ID() valueobjects.GCASProportionCostID {
	return e.id
}

// EventID - EventID のゲッター
func (e *GCASProportionCost) EventID() valueobjects.EventID {
	return e.eventID
}

// AccountID - AccountID のゲッター
func (e *GCASProportionCost) AccountID() string {
	return e.accountID
}

// Data - Data のゲッター
func (e *GCASProportionCost) Data() *GCASProportionCostData {
	return e.data
}

// Meta - Meta のゲッター
func (e *GCASProportionCost) Meta() *Meta {
	return e.meta
}
