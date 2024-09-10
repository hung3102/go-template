package authentication

import "golang.org/x/xerrors"

// ErrAuthTokenExpired - AuthTokenが切れているときに返すエラー
var ErrAuthTokenExpired = xerrors.New("auth token expired")
