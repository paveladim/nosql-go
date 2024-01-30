package endpoints

import (
	"net/http"
	"nosql-go/internal/mongodb"

	"github.com/gin-gonic/gin"
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
