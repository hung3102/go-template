package repositoryerrors

import "golang.org/x/xerrors"

// NestedTxError - トランザクションがネストされているときに返すエラー
type NestedTxError struct{}

// NewNestedTxError - NestedTxError のコンストラクタ
func NewNestedTxError(message string) RepositoryError[NestedTxError] {
	return &repositoryError[NestedTxError]{
		msg:   message,
		frame: xerrors.Caller(1),
	}
}
