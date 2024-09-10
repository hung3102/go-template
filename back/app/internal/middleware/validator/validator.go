// Package validator - リクエストを検証するためのミドルウェア
//
// https://github.com/deepmap/oapi-codegen/blob/master/pkg/middleware/oapi_validate.go
// 最初はこのファイルを使用していたが、エラーを勝手に破壊してきて辛かったので、一部を切り出して利用する。
package validator

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type (
	// EchoContextKey - is a key for echo.Context.
	EchoContextKey = struct{}
)

// Middleware - middleware for validating requests.
type Middleware struct {
	router *routers.Router
}

// NewMiddleware - creates a new Middleware.
func NewMiddleware(specs *openapi3.T) (*Middleware, error) {
	router, err := gorillamux.NewRouter(specs)
	if err != nil {
		return nil, xerrors.Errorf("failed to create router: %w", err)
	}

	vm := &Middleware{
		router: &router,
	}

	return vm, nil
}

// Middleware - middleware for validating requests.
func (vm *Middleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := c.Request()
		route, pathParams, err := (*vm.router).FindRoute(req)

		if err != nil {
			switch e := err.(type) {
			case *routers.RouteError:
				return xerrors.Errorf("failed to find route: %w", e)
			default:
				return xerrors.Errorf("failed to validate route: %w", err)
			}
		}

		validationInput := &openapi3filter.RequestValidationInput{
			Request:    req,
			PathParams: pathParams,
			Route:      route,
			Options: &openapi3filter.Options{
				MultiError: true,
			},
		}
		requestContext := context.WithValue(ctx, EchoContextKey{}, ctx)

		err = openapi3filter.ValidateRequest(requestContext, validationInput)
		if err != nil {
			return xerrors.Errorf("failed to validate request: %w", err)
		}

		return next(c)
	}
}
