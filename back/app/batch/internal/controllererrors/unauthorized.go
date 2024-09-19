package controllererrors

import (
	"net/http"

	"github.com/topgate/gcim-temporary/back/app/batch/internal/interfaces/openapi"
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// UnauthorizedError - 認証に失敗したとき返すエラー
type UnauthorizedError struct{}

// JSON - JSONっぽいオブジェクトを返す
func (n UnauthorizedError) JSON(errorCode, message string) any {
	return openapi.ErrorUnauthorized{
		ErrorCode: errorCode,
		Reason:    message,
		Kind:      openapi.Unauthorized,
	}
}

// NewUnauthorizedError - constructor of UnauthorizedError
func NewUnauthorizedError(errorCode code.ErrorCode, message string, err error) ControllerError[UnauthorizedError] {
	return &controllerError[UnauthorizedError]{
		code:       errorCode,
		statusCode: http.StatusUnauthorized,
		msg:        message,
		err:        err,
		frame:      xerrors.Caller(1),
		typ:        UnauthorizedError{},
	}
}
