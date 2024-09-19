package controllererrors

import (
	"fmt"

	code "github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"golang.org/x/xerrors"
)

type multiError[T JSON] struct {
	code       code.ErrorCode
	statusCode int
	msg        string
	err        error
	frame      xerrors.Frame
	typ        T
}

func (m multiError[T]) Error() string {
	//TODO implement me
	panic("implement me")
}

func (m multiError[T]) FormatError(p xerrors.Printer) (next error) {
	//TODO implement me
	panic("implement me")
}

func (m multiError[T]) Format(f fmt.State, verb rune) {
	//TODO implement me
	panic("implement me")
}

func (m multiError[T]) StatusCode() int {
	//TODO implement me
	panic("implement me")
}

func (m multiError[T]) Code() code.ErrorCode {
	//TODO implement me
	panic("implement me")
}

func (m multiError[T]) JSON() any {
	//TODO implement me
	panic("implement me")
}
