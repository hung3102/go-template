package valueobjects

import "github.com/rs/xid"

// OrgCSPAccountCostID - 団体ごとのCSPごとのアカウントごとID
type OrgCSPAccountCostID xid.ID

// NewOrgCSPAccountCostID - 新しいOrgCSPAccountCostIDを生成する
func NewOrgCSPAccountCostID() OrgCSPAccountCostID {
	return OrgCSPAccountCostID(xid.New())
}

// NewOrgCSPAccountCostIDFromString - 文字列をOrgCSPAccountCostIDに変換する
func NewOrgCSPAccountCostIDFromString(str string) (OrgCSPAccountCostID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return OrgCSPAccountCostID(xid.NilID()), err
	}
	return OrgCSPAccountCostID(x), nil
}

// String - OrgCSPAccountCostIDを文字列に変換する
func (v OrgCSPAccountCostID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - OrgCSPAccountCostIDを文字列に変換する
func (v OrgCSPAccountCostID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
