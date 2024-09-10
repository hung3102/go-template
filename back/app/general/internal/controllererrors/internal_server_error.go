package controllererrors

import (
	"net/http"

	"github.com/topgate/gcim-temporary/back/app/general/internal/interfaces/openapi"
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// InternalServerError - サーバ内部エラー
type InternalServerError struct{}

// JSON - JSONっぽいオブジェクトを返す
func (n InternalServerError) JSON(errorCode, message string) any {
	return openapi.ErrorInternalServerError{
		ErrorCode: errorCode,
		Reason:    message,
		Kind:      openapi.InternalServerError,
	}
}

// NewInternalServerError - constructor of InternalServerError
func NewInternalServerError(errorCode code.ErrorCode, message string, err error) ControllerError[InternalServerError] {
	return &controllerError[InternalServerError]{
		code:       errorCode,
		statusCode: http.StatusInternalServerError,
		msg:        message,
		err:        err,
		frame:      xerrors.Caller(1),
		typ:        InternalServerError{},
	}
}
