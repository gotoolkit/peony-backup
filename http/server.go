package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
	"github.com/gotoolkit/peony/http/handler"
	"github.com/gotoolkit/peony/http/middleware"
)

// Server implements the peony.Server interface
type Server struct {
	BindAddress string
	Debug       bool
	UserService peony.UserService
	JWTService  peony.JWTService
	Handler     *handler.Handler
}

// Start starts the HTTP server
func (server *Server) Start() error {
	if !server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	middlewareService := &middleware.Service{
		JWTService: server.JWTService,
	}

	engine := gin.New()

	engine.Use(gin.Logger(), gin.Recovery())

	dockerHandler := handler.NewDockerHandler()
	userHandler := handler.NewUserHandler(middlewareService)
	userHandler.UserService = server.UserService

	server.Handler = &handler.Handler{
		DockerHandler: dockerHandler,
		UserHandler:   userHandler,
	}

	server.Handler.SetupRoutes(engine)

	return engine.Run(server.BindAddress)

}
