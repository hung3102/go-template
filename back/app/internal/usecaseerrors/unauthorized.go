package usecaseerrors

import (
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// UnauthorizedError - 認証に失敗したとき返すエラー
type UnauthorizedError struct{}

// NewUnauthorizedError - constructor of UnauthorizedError
func NewUnauthorizedError(errorCode code.ErrorCode, message string, err error) UseCaseError[UnauthorizedError] {
	return &useCaseError[UnauthorizedError]{
		code:  errorCode,
		msg:   message,
		err:   err,
		frame: xerrors.Caller(1),
		typ:   UnauthorizedError{},
	}
}
