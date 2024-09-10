package repositoryerrors

import "golang.org/x/xerrors"

// DuplicateError - 重複エンティティを指定したときに返すエラー
type DuplicateError struct{}

// NewDuplicateError - DuplicateError のコンストラクタ
func NewDuplicateError(err error) RepositoryError[DuplicateError] {
	return &repositoryError[DuplicateError]{
		msg:   "duplicate",
		err:   err,
		frame: xerrors.Caller(1),
	}
}
