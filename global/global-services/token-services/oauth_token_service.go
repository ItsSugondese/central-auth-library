package token_services

import (
	authentication_middleware "github.com/ItsSugondese/central-auth-library-go/pkg/middleware/authentication-middleware"
	oauth_token "github.com/ItsSugondese/central-auth-library-go/pkg/utils/token/oauth-token"
	"github.com/gin-gonic/gin"
)

type OauthTokenService struct {
	maker oauth_token.OAuthMaker
}

func NewOauthTokenService(oAuthMaker oauth_token.OAuthMaker) *OauthTokenService {
	return &OauthTokenService{maker: oAuthMaker}
}

func (o *OauthTokenService) AuthMiddleware() gin.HandlerFunc {
	return authentication_middleware.OauthMiddleware(o.maker)
}

func (o *OauthTokenService) GenerateToken(c *gin.Context, userId string) (string, error) {
	err := o.maker.CreateToken(c.Writer, c.Request)

	if err != nil {
		return "", err
	}

	return "", nil
}
