package logger

import (
	"github.com/labstack/echo/v4"
	log "github.com/topgate/gcim-temporary/back/app/internal/logger"
	"go.uber.org/zap"
)

// InjectionMiddleware - loggerをcontextに注入するmiddleware
func InjectionMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := log.Inject(c.Request().Context(), logger)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
