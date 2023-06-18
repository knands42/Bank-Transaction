package jwt

import (
	"errors"
	"fmt"
	"time"

	token_maker "github.com/caiofernandes00/Bank-Transaction.git/app/internal/token"
	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (token_maker.Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (*token_maker.Payload, string, error) {
	payload, err := token_maker.NewPayload(username, duration)
	if err != nil {
		return nil, "", err
	}

	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(maker.secretKey))
	return payload, jwtToken, err
}

func (maker *JWTMaker) VerifyToken(token string) (*token_maker.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, token_maker.ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &token_maker.Payload{}, keyFunc)
	if err != nil {
		validation, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validation.Inner, token_maker.ErrExpiredToken) {
			return nil, token_maker.ErrExpiredToken
		}
		return nil, token_maker.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*token_maker.Payload)
	if !ok {
		return nil, token_maker.ErrInvalidToken
	}

	return payload, nil
}
