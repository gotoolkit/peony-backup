package handler

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
	httperror "github.com/gotoolkit/peony/http/error"
	"github.com/gotoolkit/peony/http/security"
)

const (
	// ErrInvalidCredentialsFormat is an error raised when credentials format is not valid
	ErrInvalidCredentialsFormat = peony.Error("Invalid credentials format")
	// ErrInvalidJSON defines an error raised the app is unable to parse request data
	ErrInvalidJSON = peony.Error("Invalid JSON")
	// ErrAuthDisabled is an error raised when trying to access the authentication endpoints
	// when the server has been started with the --no-auth flag
	ErrAuthDisabled = peony.Error("Authentication is disabled")
)

// AuthHandler represents an HTTP API handler for managing authentication.
type AuthHandler struct {
	authDisabled bool
	UserService  peony.UserService
	JWTService   peony.JWTService
}

// NewAuthHandler returns a new instance of AuthHandler.
func NewAuthHandler(router gin.IRouter, bouncer *security.RequestBouncer, authDisabled bool) *AuthHandler {
	h := &AuthHandler{
		authDisabled: authDisabled,
	}
	router.POST("/auth", bouncer.PublicAccess(), h.handleAuth)

	return h
}

type postAuthRequest struct {
	Username string `valid:"required"`
	Password string `valid:"required"`
}

func (handler *AuthHandler) handleAuth(c *gin.Context) {
	if handler.authDisabled {
		httperror.WriteErrorResponse(c, ErrAuthDisabled, http.StatusServiceUnavailable)
		return
	}

	var req postAuthRequest
	if err := c.BindJSON(&req); err != nil {
		httperror.WriteErrorResponse(c, ErrInvalidJSON, http.StatusBadRequest)
		return
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		httperror.WriteErrorResponse(c, ErrInvalidCredentialsFormat, http.StatusBadRequest)
		return
	}
	username := req.Username

	u, err := handler.UserService.UserByUsername(username)
	if err != nil {
		httperror.WriteErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	tokenData := &peony.TokenData{
		ID:       u.ID,
		Username: u.Username,
	}

	token, err := handler.JWTService.GenerateToken(tokenData)
	if err != nil {
		httperror.WriteErrorResponse(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jwt": token,
	})
	// if jUsername == "manu" && json.Password == "123" {
	// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// } else {

}
