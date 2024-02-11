package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"nosql-go/internal/models"
	"nosql-go/internal/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type Endpoints struct {
	Logger   *zap.Logger
	Database *mongodb.Mongo
}

func NewEndpoints(logger *zap.Logger, db *mongodb.Mongo) *Endpoints {
	return &Endpoints{
		Logger:   logger,
		Database: db,
	}
}

func (ep *Endpoints) Hello(c *gin.Context) {
	ans := map[string]string{"hello": "world"}
	c.JSON(http.StatusOK, ans)
}

func (ep *Endpoints) GetAllClients(c *gin.Context) {
	fmt.Println("GET ALL CLIENTS")
	collection := ep.Database.Database.Database("app").Collection("clients")
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Errorf("failed to get clients %w\n", err)
		fmt.Printf("failed to get clients %w\n", err)
		fmt.Println("FAILED TO GET CLIENTS")
	}

	var ans []models.Client
	for cursor.Next(context.TODO()) {
		var client models.ClientMg
		err := cursor.Decode(&client)

		if err != nil {
			fmt.Errorf("failed to get client %w\n", err)
			return
		}

		ans = append(ans, *models.MapClient(&client))
	}

	c.JSON(http.StatusOK, ans)
}
