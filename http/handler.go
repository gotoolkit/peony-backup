package http

import (
	"github.com/gin-gonic/gin"
)

// Handler is a collection of all the service handlers.
type Handler struct {
	DockerHandler *DockerHandler
}

// SetupRoutes setup handler
func (h *Handler) SetupRoutes(router *gin.Engine) {
	// router.Group("/docker", h.DockerHandler.Handlers...)
}
