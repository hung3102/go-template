package valueobjects

import "github.com/rs/xid"

// UserSessionID - ユーザーセッションID
type UserSessionID xid.ID

// NewUserSessionID - 新しいUserSessionIDを生成する
func NewUserSessionID() UserSessionID {
	return UserSessionID(xid.New())
}

// NewUserSessionIDFromString - 文字列をUserSessionIDに変換する
func NewUserSessionIDFromString(str string) (UserSessionID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return UserSessionID(xid.NilID()), err
	}
	return UserSessionID(x), nil
}

// String - UserSessionIDを文字列に変換する
func (v UserSessionID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - UserSessionIDの値がnilか判定する
func (v UserSessionID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
