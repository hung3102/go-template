package controllererrors

import (
	"net/http"

	"github.com/topgate/gcim-temporary/back/app/general/internal/interfaces/openapi"
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// UnavailableError - 利用不可エラー
type UnavailableError struct{}

// JSON - JSONっぽいオブジェクトを返す
func (n UnavailableError) JSON(errorCode, message string) any {
	return openapi.ErrorUnavailableError{
		ErrorCode: errorCode,
		Reason:    message,
		Kind:      openapi.Unavailable,
	}
}

// NewUnavailableError - constructor of UnavailableError
func NewUnavailableError(errorCode code.ErrorCode, message string, err error) ControllerError[UnavailableError] {
	return &controllerError[UnavailableError]{
		code:       errorCode,
		statusCode: http.StatusUnauthorized,
		msg:        message,
		err:        err,
		frame:      xerrors.Caller(1),
		typ:        UnavailableError{},
	}
}
