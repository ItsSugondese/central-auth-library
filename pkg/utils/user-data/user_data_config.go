package user_data

import (
	oauth2_setup "github.com/ItsSugondese/central-auth-library/config"
	token_services "github.com/ItsSugondese/central-auth-library/global/global-services/token-services"
	dto_utils "github.com/ItsSugondese/central-auth-library/pkg/utils/dto-utils"
	"github.com/gin-gonic/gin"
)

func GetTokenDetailsContext(ctx *gin.Context, tokenService token_services.TokenAuthService) (*oauth2_setup.JWTCustomAccessClaims, error) {
	claimsMap, err := tokenService.DecryptTokenContext(ctx)
	if err != nil {
		return nil, err
	}

	claimsJwt := &oauth2_setup.JWTCustomAccessClaims{}
	err = dto_utils.DtoConvertErrorHandledReturnError(claimsMap, &claimsJwt)

	return claimsJwt, err
}
