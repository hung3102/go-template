package repositoryerrors

import "golang.org/x/xerrors"

// UnknownError - 不明なエラー
type UnknownError struct{}

// NewUnknownError - UnknownError のコンストラクタ
func NewUnknownError(message string, err error) RepositoryError[UnknownError] {
	return &repositoryError[UnknownError]{
		msg:   message,
		err:   err,
		frame: xerrors.Caller(1),
	}
}
