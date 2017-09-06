package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DockerHandler represents an HTTP API handler for proxying requests to the Docker API.
type DockerHandler struct {
}

// NewDockerHandler returns a new instance of DockerHandler.
func NewDockerHandler() *DockerHandler {
	h := &DockerHandler{}

	return h
}

// Routes setup docker handler routers
func (h *DockerHandler) Routes(router *gin.RouterGroup) {

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})

}
