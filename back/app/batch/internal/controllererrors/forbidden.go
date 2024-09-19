package controllererrors

import (
	"net/http"

	"github.com/topgate/gcim-temporary/back/app/batch/internal/interfaces/openapi"
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// ForbiddenError - 許可されないエンティティを取得しようとしたときに返すエラー
type ForbiddenError struct{}

// JSON - JSONっぽいオブジェクトを返す
func (n ForbiddenError) JSON(errorCode, message string) any {
	return openapi.ErrorForbidden{
		ErrorCode: errorCode,
		Reason:    message,
		Kind:      openapi.Forbidden,
	}
}

// NewForbiddenError - constructor of ForbiddenError
func NewForbiddenError(errorCode code.ErrorCode, message string, err error) ControllerError[ForbiddenError] {
	return &controllerError[ForbiddenError]{
		code:       errorCode,
		statusCode: http.StatusForbidden,
		msg:        message,
		err:        err,
		frame:      xerrors.Caller(1),
		typ:        ForbiddenError{},
	}
}
