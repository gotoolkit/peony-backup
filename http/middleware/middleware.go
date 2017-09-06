package middleware

import (
	"github.com/gotoolkit/peony"
)

// MiddlewareService represents a service to manage HTTP middlewares
type MiddlewareService struct {
	JWTService peony.JWTService
}
