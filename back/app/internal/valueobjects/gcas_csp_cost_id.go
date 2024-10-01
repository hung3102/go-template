package valueobjects

import "github.com/rs/xid"

// GCASCSPCostID - GCAS Dashboardから貰ったCSPのトータルコストID
type GCASCSPCostID xid.ID

// NewGCASCSPCostID - 新しいGCASCSPCostIDを生成する
func NewGCASCSPCostID() GCASCSPCostID {
	return GCASCSPCostID(xid.New())
}

// NewGCASCSPCostIDFromString - 文字列をGCASCSPCostIDに変換する
func NewGCASCSPCostIDFromString(str string) (GCASCSPCostID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return GCASCSPCostID(xid.NilID()), err
	}
	return GCASCSPCostID(x), nil
}

// String - GCASCSPCostIDを文字列に変換する
func (v GCASCSPCostID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - GCASCSPCostIDを文字列に変換する
func (v GCASCSPCostID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
