package paseto_token

//import (
//	"fmt"
//	"os"
//	"time"
//
//	"github.com/aead/chacha20poly1305"
//	"github.com/o1egl/paseto"
//)
//
//type PasetoMaker struct {
//	Paseto       *paseto.V2
//	SymmetricKey []byte
//	Duration     time.Duration
//}
//
//var TokenMaker *PasetoMaker
//
//func NewPasetoMaker() (*PasetoMaker, error) {
//	symmetricKey := os.Getenv("PASETO_SYMMETRIC_KEY")
//	if len(symmetricKey) != chacha20poly1305.KeySize {
//		return nil, fmt.Errorf("SymmetricKey too short should be: %v", chacha20poly1305.KeySize)
//	}
//
//	expireTime := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
//	expireDuration, err := time.ParseDuration(expireTime)
//
//	if err != nil {
//		return nil, err
//	}
//
//	maker := &PasetoMaker{
//		Paseto:       paseto.NewV2(),
//		SymmetricKey: []byte(symmetricKey),
//		Duration:     expireDuration,
//	}
//
//	return maker, nil
//}
//
//func (maker *PasetoMaker) CreateToken(userId string) (string, error) {
//	payload, err := NewPayload(userId, maker.Duration)
//	if err != nil {
//		return "", err
//	}
//
//	return maker.Paseto.Encrypt(maker.SymmetricKey, payload, nil)
//}
//
//func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
//	payload := &Payload{}
//
//	err := maker.Paseto.Decrypt(token, maker.SymmetricKey, payload, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	err = payload.Valid()
//	if err != nil {
//		return nil, err
//	}
//
//	return payload, nil
//}
