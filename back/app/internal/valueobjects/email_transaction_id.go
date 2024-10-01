package valueobjects

import "github.com/rs/xid"

// EmailTransactionID - email_transactionID
type EmailTransactionID xid.ID

// NewEmailTransactionID - 新しいEmailTransactionIDを生成する
func NewEmailTransactionID() EmailTransactionID {
	return EmailTransactionID(xid.New())
}

// NewEmailTransactionIDFromString - 文字列をEmailTransactionIDに変換する
func NewEmailTransactionIDFromString(str string) (EmailTransactionID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return EmailTransactionID(xid.NilID()), err
	}
	return EmailTransactionID(x), nil
}

// String - EmailTransactionIDを文字列に変換する
func (v EmailTransactionID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - EmailTransactionIDの値がnilか判定する
func (v EmailTransactionID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
