package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
)

// Server implements the peony.Server interface
type Server struct {
	BindAddress string
	Debug       bool
	UserService peony.UserService
	Handler     *Handler
}

// Start starts the HTTP server
func (server *Server) Start() error {
	if !server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(gin.Logger(), gin.Recovery())

	dockerHandler := NewDockerHandler(engine)
	server.Handler = &Handler{
		DockerHandler: dockerHandler,
	}

	// server.Handler.SetupRoutes(engine)

	return engine.Run(server.BindAddress)

}
