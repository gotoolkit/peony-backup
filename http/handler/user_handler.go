package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
	"github.com/gotoolkit/peony/http/middleware"
)

// UserHandler represents the User API.
type UserHandler struct {
	UserService       peony.UserService
	middlewareService *middleware.MiddlewareService
}

// NewUserHandler returns a new instance of DockerHandler.
func NewUserHandler(middlewareService *middleware.MiddlewareService) *UserHandler {
	h := &UserHandler{
		middlewareService: middlewareService,
	}
	return h
}

// Routes defines the implementation of hanlder.IHandler
func (h *UserHandler) Routes(router *gin.RouterGroup) {

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})
}
