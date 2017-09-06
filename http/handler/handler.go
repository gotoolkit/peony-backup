package handler

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Routes(*gin.RouterGroup)
}

// Handler is a collection of all the service handlers.
type Handler struct {
	DockerHandler IHandler
	UserHandler   IHandler
}

// SetupRoutes setup all the service handlers
func (h *Handler) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		h.DockerHandler.Routes(v1.Group("/docker"))
		h.UserHandler.Routes(v1.Group("/user"))
	}
}
