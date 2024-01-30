package server

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"server",
		fx.Provide(
			NewServerConfig,
			NewServer,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *Server) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						server.OnStart()
						return nil
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("server")
		}),
	)
}
