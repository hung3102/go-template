package valueobjects

import "github.com/rs/xid"

// PaymentID - 収納情報ID
type PaymentID xid.ID

// NewPaymentID - 新しいPaymentIDを生成する
func NewPaymentID() PaymentID {
	return PaymentID(xid.New())
}

// NewPaymentIDFromString - 文字列をPaymentIDに変換する
func NewPaymentIDFromString(str string) (PaymentID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return PaymentID(xid.NilID()), err
	}
	return PaymentID(x), nil
}

// String - PaymentIDを文字列に変換する
func (v PaymentID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - PaymentIDの値がnilか判定する
func (v PaymentID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
