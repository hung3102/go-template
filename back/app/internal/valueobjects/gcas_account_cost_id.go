package valueobjects

import "github.com/rs/xid"

// GCASAccountCostID - GCAS Dashboardから貰ったコスト情報ID
type GCASAccountCostID xid.ID

// NewGCASAccountCostID - 新しいGCASAccountCostIDを生成する
func NewGCASAccountCostID() GCASAccountCostID {
	return GCASAccountCostID(xid.New())
}

// NewGCASAccountCostIDFromString - 文字列をGCASAccountCostIDに変換する
func NewGCASAccountCostIDFromString(str string) (GCASAccountCostID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return GCASAccountCostID(xid.NilID()), err
	}
	return GCASAccountCostID(x), nil
}

// String - GCASAccountCostIDを文字列に変換する
func (v GCASAccountCostID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - GCASAccountCostIDを文字列に変換する
func (v GCASAccountCostID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
