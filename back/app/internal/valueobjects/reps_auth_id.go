package valueobjects

import "github.com/rs/xid"

// REPSAuthID - REPSアクセストークンID
type REPSAuthID xid.ID

// NewREPSAuthID - 新しいREPSAuthIDを生成する
func NewREPSAuthID() REPSAuthID {
	return REPSAuthID(xid.New())
}

// NewREPSAuthIDFromString - 文字列をREPSAuthIDに変換する
func NewREPSAuthIDFromString(str string) (REPSAuthID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return REPSAuthID(xid.NilID()), err
	}
	return REPSAuthID(x), nil
}

// String - REPSAuthIDを文字列に変換する
func (v REPSAuthID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - REPSAuthIDの値がnilか判定する
func (v REPSAuthID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
