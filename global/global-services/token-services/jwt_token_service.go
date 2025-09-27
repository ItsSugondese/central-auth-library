package token_services

import (
	authentication_middleware "central-auth-library/pkg/middleware/authentication-middleware"
	jwt_token "central-auth-library/pkg/utils/token/jwt-token"
	"github.com/gin-gonic/gin"
)

type JwtTokenService struct {
	maker jwt_token.JwtMaker
}

func NewJwtTokenService(jwtMaker jwt_token.JwtMaker) *JwtTokenService {
	return &JwtTokenService{maker: jwtMaker}
}

func (j *JwtTokenService) AuthMiddleware() gin.HandlerFunc {
	return authentication_middleware.JwtAuthMiddleware(j.maker)
}

func (j *JwtTokenService) GenerateToken(c *gin.Context, userId string) (string, error) {
	return j.maker.CreateToken(userId)
}
