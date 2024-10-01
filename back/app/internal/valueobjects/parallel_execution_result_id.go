package valueobjects

import "github.com/rs/xid"

// ParallelExecutionResultID - 並列実行結果ID
type ParallelExecutionResultID xid.ID

// NewParallelExecutionResultID - 新しいParallelExecutionResultIDを生成する
func NewParallelExecutionResultID() ParallelExecutionResultID {
	return ParallelExecutionResultID(xid.New())
}

// NewParallelExecutionResultIDFromString - 文字列をParallelExecutionResultIDに変換する
func NewParallelExecutionResultIDFromString(str string) (ParallelExecutionResultID, error) {
	x, err := xid.FromString(str)
	if err != nil {
		return ParallelExecutionResultID(xid.NilID()), err
	}
	return ParallelExecutionResultID(x), nil
}

// String - ParallelExecutionResultIDを文字列に変換する
func (v ParallelExecutionResultID) String() string {
	x := xid.ID(v)
	return x.String()
}

// IsNil - ParallelExecutionResultIDを文字列に変換する
func (v ParallelExecutionResultID) IsNil() bool {
	x := xid.ID(v)
	return x.IsNil()
}
