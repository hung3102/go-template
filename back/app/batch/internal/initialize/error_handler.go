package initialize

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/labstack/echo/v4"
	"github.com/topgate/gcim-temporary/back/app/batch/internal/controllererrors"
	"github.com/topgate/gcim-temporary/back/app/batch/internal/interfaces/openapi"
	"github.com/topgate/gcim-temporary/back/pkg/environ"
)

const serverError = "server error"

// ErrorHandler - error handler
func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		logger := c.Logger()
		var (
			httpErr  *echo.HTTPError
			routeErr *routers.RouteError
			multiErr openapi3.MultiError
		)

		// RouteErrorがあればRouteErrorとして返す
		if errors.As(err, &routeErr) {
			ce := &openapi.ErrorInternalServerError{
				Reason: "",
			}

			switch {
			case errors.Is(routeErr, routers.ErrPathNotFound):
				ce.Reason = "path not found"
				_ = c.JSON(http.StatusNotFound, ce)
			case errors.Is(routeErr, routers.ErrMethodNotAllowed):
				ce.Reason = "method not allowed"
				_ = c.JSON(http.StatusMethodNotAllowed, ce)
			default:
				logger.Warnf("unknown route error: %v", routeErr)
				ce.Reason = serverError
				_ = c.JSON(http.StatusInternalServerError, ce)
			}
			return
		}

		// MultiErrorがあればMultiErrorとして返す
		// MultiErrorはOpenAPIのバリデーションエラーのはず
		if errors.As(err, &multiErr) {
			returnMultiError := make([]openapi.ErrorValidationError, 0, len(multiErr))
			for _, e := range multiErr {
				var (
					schemaErr *openapi3.SchemaError
				)
				ce := &openapi.ErrorInternalServerError{
					Reason: "",
				}

				if errors.As(e, &schemaErr) {
					// map[string]any への変換が大変だったため一旦JSONを経由してる。余裕ができたら(たぶんない)削除する。
					jb, err := json.Marshal(schemaErr.Schema)
					if err != nil {
						logger.Errorf("Failed to marshal scheme: %v", err)
						ce.Reason = serverError
						_ = c.JSON(http.StatusInternalServerError, ce)
						return
					}
					s := make(map[string]any)
					err = json.Unmarshal(jb, &s)
					if err != nil {
						logger.Errorf("Failed to unmarshal scheme: %v", err)
						ce.Reason = serverError
						_ = c.JSON(http.StatusInternalServerError, ce)
						return
					}

					if !environ.IsLocal() {
						s["Description"] = ""
					}

					returnMultiError = append(returnMultiError, openapi.ErrorValidationError{
						Reason: schemaErr.Reason,
						Value:  schemaErr.Value,
						Rule:   schemaErr.SchemaField,
						Schema: s,
					})
				}
			}

			_ = c.JSON(http.StatusBadRequest, returnMultiError)
			return
		}

		var cerr controllererrors.ControllerError[controllererrors.JSON]
		if errors.As(err, &cerr) {
			_ = c.JSON(cerr.StatusCode(), cerr.JSON())
			return
		}

		// HTTP errorがあったらHTTP errorとして返す
		if errors.As(err, &httpErr) {
			_ = c.JSON(httpErr.Code, httpErr.Message)
			return
		}

		_ = c.JSON(http.StatusInternalServerError, &openapi.ErrorInternalServerError{
			Reason: "server error",
		})
	}
}
