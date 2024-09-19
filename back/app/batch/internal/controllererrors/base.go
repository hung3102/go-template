package controllererrors

import (
	"fmt"

	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"github.com/topgate/gcim-temporary/back/pkg/httpstatuscode"
	"golang.org/x/xerrors"
)

// JSON - JSONにして返すことの出来るオブジェクト
type JSON interface {
	JSON(errorCode, message string) any
}

// ControllerError - Controllerが返すエラー
type ControllerError[T JSON] interface {
	error
	xerrors.Formatter
	fmt.Formatter
	httpstatuscode.Interface
	Code() code.ErrorCode
	JSON() any
}

type controllerError[T JSON] struct {
	code       code.ErrorCode
	statusCode int
	msg        string
	err        error
	frame      xerrors.Frame
	typ        T
}

// JSON - JSONにして返すことの出来るオブジェクト
func (e *controllerError[T]) JSON() any {
	return e.typ.JSON(string(e.code), e.msg)
}

// Code - エラーコードを返す
func (e *controllerError[T]) Code() code.ErrorCode {
	return e.code
}

// StatusCode - ステータスコードを返す
func (e *controllerError[T]) StatusCode() int {
	return e.statusCode
}

// Error - エラーメッセージを返す
func (e *controllerError[T]) Error() string {
	return fmt.Sprint(e)
}

// Format - エラーメッセージをフォーマットして返す
func (e *controllerError[T]) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

// FormatError - エラーメッセージをフォーマットして返す
func (e *controllerError[T]) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

// Unwrap - エラーを返す
func (e *controllerError[T]) Unwrap() error {
	return e.err
}
