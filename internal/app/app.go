package app

import (
	"nosql-go/internal/endpoints"
	"nosql-go/internal/mongodb"
	"nosql-go/internal/router"
	"nosql-go/internal/server"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		// Наше приложение будет запускаться с помощью нескольких модулей
		fx.Options(
			server.NewModule(),
			mongodb.NewModule(),
			router.NewModule(),
			endpoints.NewModule(),
		// Какие модули будем возвращать (базы данных? телеграм)
		),
		// Реализуем какой-то конструктор для нашего конфига
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
		// Всё наше приложение будет писаться зап логгером
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
