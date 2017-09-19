package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
	"github.com/gotoolkit/peony/http/security"
)

// UserHandler represents the User API.
type UserHandler struct {
	UserService peony.UserService
}

// NewUserHandler returns a new instance of DockerHandler.
func NewUserHandler(router gin.IRouter, bouncer *security.RequestBouncer) *UserHandler {
	h := &UserHandler{}

	router.GET("/users", bouncer.RestrictedAccess(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})

	router.POST("/users/:id", bouncer.AuthenticatedAccess(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})

	router.GET("/users/admin/init", bouncer.PublicAccess(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})
	return h
}
