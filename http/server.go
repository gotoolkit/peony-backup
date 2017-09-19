package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
	"github.com/gotoolkit/peony/http/handler"
	"github.com/gotoolkit/peony/http/security"
)

// Server implements the peony.Server interface
type Server struct {
	BindAddress  string
	Debug        bool
	AuthDisabled bool
	UserService  peony.UserService
	JWTService   peony.JWTService
	Handler      *handler.Handler
}

// Start starts the HTTP server
func (server *Server) Start() error {
	if !server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	requestBouncer := security.NewRequestBouncer(server.JWTService, server.AuthDisabled)

	engine := gin.New()

	engine.Use(gin.Logger(), gin.Recovery())

	api := engine.Group("/api")
	{
		dockerHandler := handler.NewDockerHandler()

		authHandler := handler.NewAuthHandler(api, requestBouncer, server.AuthDisabled)
		authHandler.UserService = server.UserService
		authHandler.JWTService = server.JWTService
		
		userHandler := handler.NewUserHandler(api, requestBouncer)
		userHandler.UserService = server.UserService

		server.Handler = &handler.Handler{
			DockerHandler: dockerHandler,
			UserHandler:   userHandler,
		}
	}

	return engine.Run(server.BindAddress)

}
