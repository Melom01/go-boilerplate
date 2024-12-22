package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Melom01/go-boilerplate/pkg/config"
)

func ValidateToken(firebaseService *config.FirebaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		const bearerPrefix = "Bearer "

		if !strings.HasPrefix(authHeader, bearerPrefix) || len(authHeader) <= len(bearerPrefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		token := authHeader[len(bearerPrefix):]
		decodedToken, err := firebaseService.Auth.VerifyIDToken(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user", decodedToken)
		c.Next()
	}
}
