package logger

import (
	"context"

	"github.com/topgate/gcim-temporary/back/pkg/superctx"
	"go.uber.org/zap"
)

type loggerKey *zap.Logger

// Inject - loggerをcontextに注入する
func Inject(ctx context.Context, logger *zap.Logger) context.Context {
	return superctx.WithValue[loggerKey](ctx, logger)
}

// FromContext - loggerをcontextから取得する
func FromContext(ctx context.Context) *zap.Logger {
	l, ok := superctx.Value[loggerKey](ctx)
	if !ok {
		panic("logger is not set")
	}
	return l
}
