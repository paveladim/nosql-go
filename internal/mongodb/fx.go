package mongodb

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"mongodb",
		fx.Provide(
			NewMongoConfig,
			NewMongo,
		),
		fx.Invoke(),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("mongodb")
		}),
	)
}
