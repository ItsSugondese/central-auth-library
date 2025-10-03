package token_services

import (
	"github.com/gin-gonic/gin"
)

type TokenAuthService interface {
	AuthMiddleware() gin.HandlerFunc
	GenerateToken(c *gin.Context, userID string) (string, error)
	DecryptTokenContext(ctx *gin.Context) (payload map[string]interface{}, err error)
}
