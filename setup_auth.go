package central_auth_library

import (
	tokenauthtypeenums "central-auth-library/enums/struct-enums/token-auth-type-enums"
	tokenservices "central-auth-library/global/global-services/token-services"
	jwttoken "central-auth-library/pkg/utils/token/jwt-token"
	oauthtoken "central-auth-library/pkg/utils/token/oauth-token"
	"github.com/go-oauth2/oauth2/v4/server"
)

func SetupAuthToken(authType string, server *server.Server) tokenservices.TokenAuthService {
	if authType == tokenauthtypeenums.TokenAuthType.OAUTH {
		tokenMaker, err := oauthtoken.NewOAuthMaker(server)
		if err != nil {
			panic("Error Setting up token maker for OAuth: " + (err.Error()))
		} else if tokenMaker == nil {
			panic("Token Maker set to null for OAuth")
		}
		return tokenservices.NewOauthTokenService(*tokenMaker)
	} else {
		//else if authType == token_auth_type_enums.TokenAuthType.PASETO {
		//	tokenMaker, err := paseto_token.NewPasetoMaker()
		//	if err != nil {
		//		panic("Error Setting up token maker for Paseto " + err.Error())
		//	} else if tokenMaker == nil {
		//		panic("Token Maker set to null for Paseto")
		//	}
		//
		//	return token_services.NewPasetoTokenService(*tokenMaker)
		//}

		tokenMaker, err := jwttoken.NewJwtMaker()
		if err != nil {
			panic("Error Setting up token maker for OAuth: " + (err.Error()))
		} else if tokenMaker == nil {
			panic("Token Maker set to null for JWT")
		}
		return tokenservices.NewJwtTokenService(*tokenMaker)
	}
}
