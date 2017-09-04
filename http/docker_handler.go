package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DockerHandler represents an HTTP API handler for proxying requests to the Docker API.
type DockerHandler struct {
	*gin.RouterGroup
}

// NewDockerHandler returns a new instance of DockerHandler.
func NewDockerHandler(router *gin.Engine) *DockerHandler {
	h := &DockerHandler{
		RouterGroup: &router.RouterGroup,
	}

	h.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})

	return h
}

func (h *DockerHandler) Routes() {

}
