package usecaseerrors

import (
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// NoPermissionError - 権限がない時に返すエラー
type NoPermissionError struct{}

// NewNoPermissionError - constructor of NoPermissionError
func NewNoPermissionError(errorCode code.ErrorCode, message string) UseCaseError[NoPermissionError] {
	return &useCaseError[NoPermissionError]{
		code:  errorCode,
		msg:   message,
		err:   nil,
		frame: xerrors.Caller(1),
		typ:   NoPermissionError{},
	}
}
