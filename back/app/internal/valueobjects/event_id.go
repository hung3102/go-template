package valueobjects

import "github.com/rs/xid"

// EventID - イベントID
type EventID xid.ID

// NewEventID - 新しいEventIDを生成する
func NewEventID() EventID {
	return EventID(xid.New())
}

// NewEventIDFromString - 文字列をEventIDに変換する
func NewEventIDFromString(str string) (EventID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return EventID(xid.NilID()), err
	}
	return EventID(x), nil
}

// String - EventIDを文字列に変換する
func (v EventID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - EventIDの値がnilか判定する
func (v EventID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
