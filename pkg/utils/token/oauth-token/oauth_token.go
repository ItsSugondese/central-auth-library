package oauth_token

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
	"net/http"
)

type OAuthMaker struct {
	srv *server.Server
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
