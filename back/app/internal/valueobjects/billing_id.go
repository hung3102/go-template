package valueobjects

import "github.com/rs/xid"

// BillingID - 請求ID
type BillingID xid.ID

// NewBillingID - 新しいBillingIDを生成する
func NewBillingID() BillingID {
	return BillingID(xid.New())
}

// NewBillingIDFromString - 文字列をBillingIDに変換する
func NewBillingIDFromString(str string) (BillingID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return BillingID(xid.NilID()), err
	}
	return BillingID(x), nil
}

// String - BillingIDを文字列に変換する
func (v BillingID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - BillingIDを文字列に変換する
func (v BillingID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
