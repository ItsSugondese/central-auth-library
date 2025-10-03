package config

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTAccessClaims jwt claims
type JWTCustomAccessClaims struct {
	jwt.RegisteredClaims
	UserId     string `json:"userId"`
	TenantName string `json:"tenantName"`
}

// Valid claims verification
func (a *JWTCustomAccessClaims) Valid() error {
	if a.ExpiresAt != nil && time.Unix(a.ExpiresAt.Unix(), 0).Before(time.Now()) {
		return errors.ErrInvalidAccessToken
	}
	return nil
}

// NewJWTCustomAccessGenerate create to generate the jwt access token instance
func NewCustomJWTCustomAccessGenerate(kid string, key []byte, method jwt.SigningMethod) *JWTCustomAccessGenerate {
	return &JWTCustomAccessGenerate{
		SignedKeyID:  kid,
		SignedKey:    key,
		SignedMethod: method,
	}
}

// JWTCustomAccessGenerate generate the jwt access token
type JWTCustomAccessGenerate struct {
	SignedKeyID  string
	SignedKey    []byte
	SignedMethod jwt.SigningMethod
}

// Token based on the UUID generated token
func (a *JWTCustomAccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {

	extendableTokenInfo := data.TokenInfo.(oauth2.ExtendableTokenInfo).GetExtension()
	claims := jwt.MapClaims{
		"aud": jwt.ClaimStrings{data.Client.GetID()},
		"sub": data.UserID,
		"exp": jwt.NewNumericDate(data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn())),
	}

	for k, v := range extendableTokenInfo {
		if len(v) == 1 {
			claims[k] = v[0]
		} else {
			claims[k] = v
		}
	}

	token := jwt.NewWithClaims(a.SignedMethod, claims)
	if a.SignedKeyID != "" {
		token.Header["kid"] = a.SignedKeyID
	}
	var key interface{}
	if a.isEs() {
		v, err := jwt.ParseECPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return "", "", err
		}
		key = v
	} else if a.isRsOrPS() {
		v, err := jwt.ParseRSAPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return "", "", err
		}
		key = v
	} else if a.isHs() {
		key = a.SignedKey
	} else if a.isEd() {
		v, err := jwt.ParseEdPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return "", "", err
		}
		key = v
	} else {
		return "", "", errors.New("unsupported sign method")
	}

	access, err := token.SignedString(key)
	if err != nil {
		return "", "", err
	}
	refresh := ""

	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return access, refresh, nil
}

func (a *JWTCustomAccessGenerate) isEs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "ES")
}

func (a *JWTCustomAccessGenerate) isRsOrPS() bool {
	isRs := strings.HasPrefix(a.SignedMethod.Alg(), "RS")
	isPs := strings.HasPrefix(a.SignedMethod.Alg(), "PS")
	return isRs || isPs
}

func (a *JWTCustomAccessGenerate) isHs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "HS")
}

func (a *JWTCustomAccessGenerate) isEd() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "Ed")
}
