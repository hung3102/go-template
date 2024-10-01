package valueobjects

import "github.com/rs/xid"

// CreatorID - 請求書の発行元情報ID
type CreatorID xid.ID

// NewCreatorID - 新しいCreatorIDを生成する
func NewCreatorID() CreatorID {
	return CreatorID(xid.New())
}

// NewCreatorIDFromString - 文字列をCreatorIDに変換する
func NewCreatorIDFromString(str string) (CreatorID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return CreatorID(xid.NilID()), err
	}
	return CreatorID(x), nil
}

// String - CreatorIDを文字列に変換する
func (v CreatorID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - CreatorIDを文字列に変換する
func (v CreatorID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
