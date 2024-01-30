package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Mongo struct {
	Logger   *zap.Logger
	Config   *Config
	Database *mongo.Client
}

func NewMongo(logger *zap.Logger, config *Config) *Mongo {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Uri))

	if err != nil {
		logger.Warn(fmt.Sprintf("Cannot connect to Mongo"))
		return nil
	}

	return &Mongo{
		Logger:   logger,
		Config:   config,
		Database: client,
	}
}
