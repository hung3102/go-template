package valueobjects

import "github.com/rs/xid"

// AuthID - REPSアクセストークンID
type AuthID xid.ID

// NewAuthID - 新しいAuthIDを生成する
func NewAuthID() AuthID {
	return AuthID(xid.New())
}

// NewAuthIDFromString - 文字列をAuthIDに変換する
func NewAuthIDFromString(str string) (AuthID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return AuthID(xid.NilID()), err
	}
	return AuthID(x), nil
}

// String - AuthIDを文字列に変換する
func (v AuthID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - AuthIDを文字列に変換する
func (v AuthID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
