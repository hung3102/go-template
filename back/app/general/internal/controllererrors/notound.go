package controllererrors

import (
	"net/http"

	"github.com/topgate/gcim-temporary/back/app/general/internal/interfaces/openapi"
	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// NotFoundError - 存在しないエンティティを取得しようとしたときに返すエラー
type NotFoundError struct{}

// JSON - JSONっぽいオブジェクトを返す
func (n NotFoundError) JSON(errorCode, message string) any {
	return openapi.ErrorNotFound{
		ErrorCode: errorCode,
		Reason:    message,
		Kind:      openapi.NotFound,
	}
}

// NewNotFoundError - constructor of NotFoundError
func NewNotFoundError(errorCode code.ErrorCode, message string, err error) ControllerError[NotFoundError] {
	return &controllerError[NotFoundError]{
		code:       errorCode,
		statusCode: http.StatusNotFound,
		msg:        message,
		err:        err,
		frame:      xerrors.Caller(1),
		typ:        NotFoundError{},
	}
}
