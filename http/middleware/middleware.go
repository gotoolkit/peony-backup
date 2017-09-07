package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
	//  "github.com/satori/go.uuid"
)

// Service represents a service to manage HTTP middlewares
type Service struct {
	JWTService peony.JWTService
}

// MiddlewareHandler provides all middleware for handlers
func (service *Service) MiddlewareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		service.secureHeaders(c)
		service.authenticate(c)
		c.Next()
	}
}

// secureHeaders provides secure headers middleware for handlers
func (*Service) secureHeaders(c *gin.Context) {
	c.Writer.Header().Add("X-Content-Type-Options", "nosniff")
	c.Writer.Header().Add("X-Frame-Options", "DENY")
	// c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
}

// authenticate provides Authentication middleware for handlers
func (service *Service) authenticate(c *gin.Context) {

	var token string

	tokens, ok := c.Request.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
	}

	if token == "" {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	err := service.JWTService.VerifyToken(token)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func (service *Service) unauthorized(c *gin.Context, code int, message string) {

}
