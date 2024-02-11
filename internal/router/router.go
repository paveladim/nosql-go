package router

import (
	"nosql-go/internal/endpoints"
	"nosql-go/internal/server"
)

func RegisterRoutes(server *server.Server, ep *endpoints.Endpoints) {
	server.Server.GET("/", ep.Hello)
	server.Server.GET("/all", ep.GetAllClients)
}
