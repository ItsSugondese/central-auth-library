package oauth_token

import (
	"fmt"
	oauth2_setup "github.com/ItsSugondese/central-auth-library/config"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type OAuthMaker struct {
	srv              *server.Server
	TokenClaimsModel *oauth2_setup.JWTCustomAccessClaims
}

var TokenMaker *OAuthMaker

func NewOAuthMaker(server *server.Server) (*OAuthMaker, error) {
	maker := &OAuthMaker{
		srv: server,
	}

	return maker, nil
}

func (maker *OAuthMaker) CreateToken(w http.ResponseWriter, r *http.Request) error {
	err := maker.srv.HandleTokenRequest(w, r)
	if err != nil {
		return err
	}

	return nil
}

func (maker *OAuthMaker) VerifyToken(request *http.Request) (oauth2.TokenInfo, error) {
	bearerToken, err := maker.srv.ValidationBearerToken(request)
	if err != nil {
		return nil, err
	}

	return bearerToken, nil
}

func (maker *OAuthMaker) DecryptToken(request *http.Request, signedKey string) (payload map[string]interface{}, err error) {

	accessToken, ok := maker.srv.AccessTokenResolveHandler(request)
	if !ok {
		return nil, errors.ErrInvalidAccessToken
	}

	claimsMap := new(jwt.MapClaims)

	parsedToken, err := jwt.ParseWithClaims(accessToken, claimsMap, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.ErrInvalidAccessToken
	}

	return *claimsMap, nil
}

func (maker *OAuthMaker) SetTokenClaimsModel(tokenClaimsModel *oauth2_setup.JWTCustomAccessClaims) {
	maker.TokenClaimsModel = tokenClaimsModel
}
