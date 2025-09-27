package authentication_middleware

import (
	oauth_token "github.com/ItsSugondese/central-auth-library-go-go/pkg/utils/token/oauth-token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OauthMiddleware(maker oauth_token.OAuthMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := maker.VerifyToken(ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access Token Not Valid"})
			return
		}

		ctx.Next()
	}
}
