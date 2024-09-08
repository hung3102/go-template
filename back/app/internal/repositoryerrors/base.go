package repositoryerrors

import (
	"fmt"

	"golang.org/x/xerrors"
)

// RepositoryError - リポジトリが返すエラー
type RepositoryError[T any] interface {
	error
	xerrors.Formatter
	fmt.Formatter
}

type repositoryError[T any] struct {
	msg   string
	err   error
	frame xerrors.Frame
}

// Error - エラーメッセージを返す
func (e *repositoryError[T]) Error() string {
	return fmt.Sprint(e)
}

// Format - エラーメッセージをフォーマットして返す
func (e *repositoryError[T]) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

// FormatError - エラーメッセージをフォーマットして返す
func (e *repositoryError[T]) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

// Unwrap - エラーを返す
func (e *repositoryError[T]) Unwrap() error {
	return e.err
}
