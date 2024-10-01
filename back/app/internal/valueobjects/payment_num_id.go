package valueobjects

import "github.com/rs/xid"

// PaymentNumID - 収納番号ID
type PaymentNumID xid.ID

// NewPaymentNumID - 新しいPaymentNumIDを生成する
func NewPaymentNumID() PaymentNumID {
	return PaymentNumID(xid.New())
}

// NewPaymentNumIDFromString - 文字列をPaymentNumIDに変換する
func NewPaymentNumIDFromString(str string) (PaymentNumID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return PaymentNumID(xid.NilID()), err
	}
	return PaymentNumID(x), nil
}

// String - PaymentNumIDを文字列に変換する
func (v PaymentNumID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - PaymentNumIDの値がnilか判定する
func (v PaymentNumID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
