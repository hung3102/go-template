package valueobjects

import "github.com/rs/xid"

// OrgCSPAccountID - 団体ごとのCSPごとのアカウントごとID
type OrgCSPAccountID xid.ID

// NewOrgCSPAccountID - 新しいOrgCSPAccountIDを生成する
func NewOrgCSPAccountID() OrgCSPAccountID {
	return OrgCSPAccountID(xid.New())
}

// NewOrgCSPAccountIDFromString - 文字列をOrgCSPAccountIDに変換する
func NewOrgCSPAccountIDFromString(str string) (OrgCSPAccountID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return OrgCSPAccountID(xid.NilID()), err
	}
	return OrgCSPAccountID(x), nil
}

// String - OrgCSPAccountIDを文字列に変換する
func (v OrgCSPAccountID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - OrgCSPAccountIDを文字列に変換する
func (v OrgCSPAccountID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
