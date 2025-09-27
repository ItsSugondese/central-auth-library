package jwt_token

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type JwtMaker struct {
	privateKey     string
	publicKey      string
	expireDuration time.Duration
}

var TokenMaker *JwtMaker

func NewJwtMaker() (*JwtMaker, error) {

	expireTime := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	expireDuration, err := time.ParseDuration(expireTime)

	if err != nil {
		return nil, err
	}

	maker := &JwtMaker{
		privateKey:     os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"),
		publicKey:      os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"),
		expireDuration: expireDuration,
	}

	return maker, nil
}

func (maker *JwtMaker) CreateToken(payload interface{}) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(maker.privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return "", fmt.Errorf("create: parse keye: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(maker.expireDuration).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (maker *JwtMaker) VerifyToken(token string) (err error) {
	// Decode the base64-encoded public key
	decodedPublicKey, err := base64.StdEncoding.DecodeString(maker.publicKey)
	if err != nil {
		panic(fmt.Sprintf("could not decode public key: %v", err))
	}

	// Parse the decoded public key
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return err
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return err
	}

	// Check the paseto-token claims and validity
	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		panic("validate: invalid paseto-token")
	}

	return nil
}
