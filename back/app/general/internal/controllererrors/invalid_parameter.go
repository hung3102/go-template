package controllererrors

import (
	"net/http"

	"github.com/topgate/gcim-temporary/back/app/general/internal/interfaces/openapi"
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// InvalidParameterError - パラメータが不正な場合のエラー
type InvalidParameterError struct{}

// JSON - JSONっぽいオブジェクトを返す
func (n InvalidParameterError) JSON(errorCode, message string) any {
	return openapi.ErrorNotFound{
		ErrorCode: errorCode,
		Reason:    message,
		Kind:      openapi.NotFound,
	}
}

// NewInvalidParameterError - constructor of InvalidParameterError
func NewInvalidParameterError(errorCode code.ErrorCode, message string, err error) ControllerError[InvalidParameterError] {
	return &controllerError[InvalidParameterError]{
		code:       errorCode,
		statusCode: http.StatusBadRequest,
		msg:        message,
		err:        err,
		frame:      xerrors.Caller(1),
		typ:        InvalidParameterError{},
	}
}
