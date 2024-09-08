package usecaseerrors

import (
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// UnavailableError - 利用不可エラー
type UnavailableError struct{}

// NewUnavailableError - constructor of UnavailableError
func NewUnavailableError(errorCode code.ErrorCode, message string) UseCaseError[UnavailableError] {
	return &useCaseError[UnavailableError]{
		code:  errorCode,
		msg:   message,
		err:   nil,
		frame: xerrors.Caller(1),
		typ:   UnavailableError{},
	}
}
