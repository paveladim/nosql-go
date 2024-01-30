package router

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"router",
		fx.Provide(),
		fx.Invoke(RegisterRoutes),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("router")
		}),
	)
}
