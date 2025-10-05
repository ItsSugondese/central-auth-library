package authentication_middleware

import (
	"fmt"
	response_status_enum "github.com/ItsSugondese/central-auth-library/enums/interface-enums/response/response-status-enum"
	globaldto "github.com/ItsSugondese/central-auth-library/global/global_dto"
	oauth_token "github.com/ItsSugondese/central-auth-library/pkg/utils/token/oauth-token"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func OauthMiddleware(maker *oauth_token.OAuthMaker) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		_, err := maker.VerifyToken(ctx.Request)
//		if err != nil {
//			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access Token Not Valid"})
//			return
//		}
//
//		ctx.Next()
//	}
//}

func OauthMiddleware(maker *oauth_token.OAuthMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := maker.DecryptToken(c.Request)

		if err != nil {
			response := &globaldto.ApiResponse{
				Status:  response_status_enum.Fail(),
				Message: fmt.Sprintf("%v", err),
				Data:    []string{err.Error()},
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Next()
	}
}
