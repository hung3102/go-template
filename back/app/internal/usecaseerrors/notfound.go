package usecaseerrors

import (
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// NotFoundError - 指定されたものが存在しない時に返すエラー
type NotFoundError struct{}

// NewNotFoundError - constructor of NotFoundError
func NewNotFoundError(errorCode code.ErrorCode, message string, err error) UseCaseError[NotFoundError] {
	return &useCaseError[NotFoundError]{
		code:  errorCode,
		msg:   message,
		err:   err,
		frame: xerrors.Caller(1),
		typ:   NotFoundError{},
	}
}
