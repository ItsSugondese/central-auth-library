package token_middleware

import (
	"context"
	token_services "github.com/ItsSugondese/central-auth-library/global/global-services/token-services"
	user_data "github.com/ItsSugondese/central-auth-library/pkg/utils/user-data"
	"github.com/gin-gonic/gin"
)

func TokenClaimsMiddleware(tokenService token_services.TokenAuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenDetails, err := user_data.GetTokenDetailsContext(c, tokenService)
		if err != nil {
			return
		}

		ctx := context.WithValue(c.Request.Context(), "UserId", tokenDetails.UserId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
