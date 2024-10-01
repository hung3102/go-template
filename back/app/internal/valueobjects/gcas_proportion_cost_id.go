package valueobjects

import "github.com/rs/xid"

// GCASProportionCostID - GCASから貰った費用按分情報ID
type GCASProportionCostID xid.ID

// NewGCASProportionCostID - 新しいGCASProportionCostIDを生成する
func NewGCASProportionCostID() GCASProportionCostID {
	return GCASProportionCostID(xid.New())
}

// NewGCASProportionCostIDFromString - 文字列をGCASProportionCostIDに変換する
func NewGCASProportionCostIDFromString(str string) (GCASProportionCostID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return GCASProportionCostID(xid.NilID()), err
	}
	return GCASProportionCostID(x), nil
}

// String - GCASProportionCostIDを文字列に変換する
func (v GCASProportionCostID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - GCASProportionCostIDを文字列に変換する
func (v GCASProportionCostID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
