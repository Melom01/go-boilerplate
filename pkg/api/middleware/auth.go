package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Melom01/go-boilerplate/pkg/config"
)

func ValidateTokenMiddleware(firebaseService *config.FirebaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		token := authHeader[len("Bearer "):]
		decodedToken, err := firebaseService.Auth.VerifyIDToken(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user", decodedToken)
		c.Next()
	}
}
