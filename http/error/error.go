package error

import (
	"github.com/gin-gonic/gin"
)

// WriteErrorResponse writes an error message to the response
func WriteErrorResponse(c *gin.Context, err error, code int) {
	c.Abort()
	c.JSON(code, gin.H{
		"message": err.Error(),
	})
}
