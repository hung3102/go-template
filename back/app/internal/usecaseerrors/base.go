package usecaseerrors

import (
	"fmt"

	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

// UseCaseError - ユースケースが返すエラー
type UseCaseError[T any] interface {
	error
	xerrors.Formatter
	fmt.Formatter
	Code() code.ErrorCode
	Typ() T
}

type useCaseError[T any] struct {
	code  code.ErrorCode
	msg   string
	err   error
	frame xerrors.Frame
	typ   T
}

// Typ - エラーの型を返す
func (e *useCaseError[T]) Typ() T {
	return e.typ
}

// Code - エラーコードを返す
func (e *useCaseError[T]) Code() code.ErrorCode {
	return e.code
}

// Error - エラーメッセージを返す
func (e *useCaseError[T]) Error() string {
	return fmt.Sprint(e)
}

// Format - エラーメッセージをフォーマットして返す
func (e *useCaseError[T]) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

// FormatError - エラーメッセージをフォーマットして返す
func (e *useCaseError[T]) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

// Unwrap - Unwrap したエラーを返す
func (e *useCaseError[T]) Unwrap() error {
	return e.err
}
