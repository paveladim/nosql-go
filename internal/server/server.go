package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
	Server *gin.Engine
	Config *Config
}

func NewServer(logger *zap.Logger, config *Config) *Server {
	router := gin.Default()

	if router == nil {
		logger.Warn(fmt.Sprintf("Cannot create server"))
		return nil
	}

	return &Server{
		Logger: logger,
		Server: router,
		Config: config,
	}
}

func (server *Server) OnStart() {
	go server.Server.Run("localhost" + server.Config.Port)
}
