package usecaseerrors

import (
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// InvalidParameterError - 不正な値だったときに返すエラー
type InvalidParameterError struct{}

// NewInvalidParameterError - constructor of InvalidParameterError
func NewInvalidParameterError(errorCode code.ErrorCode, message string, err error) UseCaseError[InvalidParameterError] {
	return &useCaseError[InvalidParameterError]{
		code:  errorCode,
		msg:   message,
		err:   err,
		frame: xerrors.Caller(1),
		typ:   InvalidParameterError{},
	}
}
