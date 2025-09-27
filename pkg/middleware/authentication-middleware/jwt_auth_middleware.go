package authentication_middleware

import (
	jwt_token "central-auth-library/pkg/utils/token/jwt-token"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware is a middleware function that checks for the presence of a valid token.
func JwtAuthMiddleware(maker jwt_token.JwtMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Here you can add your authentication logic, e.g., checking for a token in the request header
		token := c.GetHeader("Authorization")

		err := maker.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
			c.Abort()
			return
		}

		// If the token is valid, proceed to the next handler
		c.Next()
	}
}
