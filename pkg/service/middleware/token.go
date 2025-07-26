package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bandanascripts/tondru/pkg/client/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractClaim(c *gin.Context) (jwt.MapClaims, error) {

	userCtx, exists := c.Get("userClaimKey")

	if !exists {
		return nil, fmt.Errorf("failed to get user claim from context")
	}

	userClaim, ok := userCtx.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("interface does not contain user claim")
	}

	return userClaim, nil
}

func AuthToken(c *gin.Context) (string, error) {

	var authHeader = c.GetHeader("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("token missing")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid token format")
	}

	var authToken = strings.TrimPrefix(authHeader, "Bearer ")

	return authToken, nil
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken, err := AuthToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userClaim, err := token.ValidateToken(c.Request.Context(), authToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("userClaimKey", userClaim)
		c.Next()
	}
}
