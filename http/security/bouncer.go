package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/peony"
	httperror "github.com/gotoolkit/peony/http/error"
)

type (
	// RequestBouncer represents an entity that manages API request accesses
	RequestBouncer struct {
		jwtService   peony.JWTService
		authDisabled bool
	}

	// RestrictedRequestContext is a data structure containing information
	// used in RestrictedAccess
	RestrictedRequestContext struct {
		IsAdmin      bool
		IsTeamLeader bool
		UserID       peony.UserID
	}
)

// PublicAccess defines a security check for public endpoints.
// No authentication is required to access these endpoints.
func (bouncer *RequestBouncer) PublicAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		mwSecureHeaders(c)
		c.Next()
	}
}

// NewRequestBouncer initializes a new RequestBouncer
func NewRequestBouncer(jwtService peony.JWTService, authDisabled bool) *RequestBouncer {
	return &RequestBouncer{
		jwtService:   jwtService,
		authDisabled: authDisabled,
	}
}

// RestrictedAccess defines a security check for restricted endpoints.
// Authentication is required to access these endpoints.
// The request context will be enhanced with a RestrictedRequestContext object
// that might be used later to authorize/filter access to resources.
func (bouncer *RequestBouncer) RestrictedAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		bouncer.mwUpgradeToRestrictedRequest(c)
		bouncer.AuthenticatedAccess()
		c.Next()
	}
}

// mwUpgradeToRestrictedRequest will enhance the current request with
// a new RestrictedRequestContext object.
func (bouncer *RequestBouncer) mwUpgradeToRestrictedRequest(c *gin.Context) {
	tokenData, err := RetrieveTokenData(c.Request)
	if err != nil {
		httperror.WriteErrorResponse(c, err, http.StatusForbidden)
		return
	}

	requestContext, err := bouncer.newRestrictedContextRequest(tokenData.ID)
	if err != nil {
		httperror.WriteErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	ctx := storeRestrictedRequestContext(c.Request, requestContext)
	c.Request.WithContext(ctx)

}

// AuthenticatedAccess defines a security check for private endpoints.
// Authentication is required to access these endpoints.
func (bouncer *RequestBouncer) AuthenticatedAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		bouncer.mwCheckAuthentication(c)
		mwSecureHeaders(c)
		c.Next()
	}
}

// mwSecureHeaders provides secure headers middleware for handlers.
func mwSecureHeaders(c *gin.Context) {
	c.Writer.Header().Add("X-Content-Type-Options", "nosniff")
	c.Writer.Header().Add("X-Frame-Options", "DENY")
}

// mwCheckAuthentication provides Authentication middleware for handlers
func (bouncer *RequestBouncer) mwCheckAuthentication(c *gin.Context) {
	var tokenData *peony.TokenData
	if !bouncer.authDisabled {
		var token string

		// Get token from the Authorization header
		tokens, ok := c.Request.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}

		if token == "" {
			httperror.WriteErrorResponse(c, peony.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		var err error
		tokenData, err = bouncer.jwtService.ParseAndVerifyToken(token)
		if err != nil {
			httperror.WriteErrorResponse(c, err, http.StatusUnauthorized)
			return
		}
	} else {
		tokenData = &peony.TokenData{}
	}
	ctx := storeTokenData(c.Request, tokenData)
	c.Request.WithContext(ctx)
}

func (bouncer *RequestBouncer) newRestrictedContextRequest(userID peony.UserID) (*RestrictedRequestContext, error) {
	requestContext := &RestrictedRequestContext{
		IsAdmin: true,
		UserID:  userID,
	}

	return requestContext, nil
}
