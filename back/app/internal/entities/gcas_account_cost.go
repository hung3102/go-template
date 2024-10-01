package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// GCASAccountCost - GCAS Dashboardから貰ったコスト情報
type GCASAccountCost struct {
	id        valueobjects.GCASAccountCostID // ID
	eventID   valueobjects.EventID           // イベントID
	accountID string                         // アカウントID
	data      *GCASAccountCostData           // もらったデータ
	meta      *Meta                          // メタ
}

// NewGCASAccountCostParam - GCAS Dashboardから貰ったコスト情報作成パラメータ
type NewGCASAccountCostParam struct {
	ID        valueobjects.GCASAccountCostID // ID
	EventID   valueobjects.EventID           // イベントID
	AccountID string                         // アカウントID
	Data      *GCASAccountCostData           // もらったデータ
	Meta      *Meta                          // メタ
}

// NewGCASAccountCost - GCAS Dashboardから貰ったコスト情報作成
func NewGCASAccountCost(param *NewGCASAccountCostParam) *GCASAccountCost {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewGCASAccountCostID()
	}
	return &GCASAccountCost{
		id:        id,
		eventID:   param.EventID,
		accountID: param.AccountID,
		data:      param.Data,
		meta:      param.Meta,
	}
}

// ID - ID のゲッター
func (e *GCASAccountCost) ID() valueobjects.GCASAccountCostID {
	return e.id
}

// EventID - EventID のゲッター
func (e *GCASAccountCost) EventID() valueobjects.EventID {
	return e.eventID
}

// AccountID - AccountID のゲッター
func (e *GCASAccountCost) AccountID() string {
	return e.accountID
}

// Data - Data のゲッター
func (e *GCASAccountCost) Data() *GCASAccountCostData {
	return e.data
}

// Meta - Meta のゲッター
func (e *GCASAccountCost) Meta() *Meta {
	return e.meta
}
