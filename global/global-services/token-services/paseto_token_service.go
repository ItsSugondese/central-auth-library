package token_services

//import (
//	authentication_middleware "github.com/ItsSugondese/central-auth-library/pkg/middleware/authentication-middleware"
//	paseto_token "github.com/ItsSugondese/central-auth-library/pkg/utils/token/paseto-token"
//	"github.com/gin-gonic/gin"
//)
//
//type PasetoTokenService struct {
//	maker paseto_token.PasetoMaker
//}
//
//func NewPasetoTokenService(pasetoMaker paseto_token.PasetoMaker) *PasetoTokenService {
//	return &PasetoTokenService{maker: pasetoMaker}
//}
//
//func (p *PasetoTokenService) AuthMiddleware() gin.HandlerFunc {
//	return authentication_middleware.PasetoAuthMiddleware(p.maker)
//}
//
//func (p *PasetoTokenService) GenerateToken(c *gin.Context, userId string) (string, error) {
//	return p.maker.CreateToken(userId)
//}
//
//func (p *PasetoTokenService) DecryptTokenContext(ctx *gin.Context) (payload map[string]interface{}, err error) {
//	return payload, err
//}
