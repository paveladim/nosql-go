package endpoints

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"endpoints",
		fx.Provide(
			NewEndpoints,
		),
		fx.Invoke(),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("endpoints")
		}),
	)
}
