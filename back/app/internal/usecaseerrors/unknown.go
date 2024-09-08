package usecaseerrors

import (
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// UnknownError - 不明なエラーが発生したときに返すエラー
type UnknownError struct{}

// NewUnknownError - constructor of UnknownError
func NewUnknownError(errorCode code.ErrorCode, message string, err error) UseCaseError[UnknownError] {
	return &useCaseError[UnknownError]{
		code:  errorCode,
		msg:   message,
		err:   err,
		frame: xerrors.Caller(1),
		typ:   UnknownError{},
	}
}
