package repositoryerrors

import "golang.org/x/xerrors"

// NotFoundError - 存在しないエンティティを指定したときに返すエラー
type NotFoundError struct{}

// NewNotFoundError - NotFoundError のコンストラクタ
func NewNotFoundError(err error) RepositoryError[NotFoundError] {
	return &repositoryError[NotFoundError]{
		msg:   "not found",
		err:   err,
		frame: xerrors.Caller(1),
	}
}
