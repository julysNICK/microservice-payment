package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < 10 {
		return nil, fmt.Errorf("invalid key size")
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)

	if err != nil {
		return "", fmt.Errorf("cannot create payload: %w", err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(maker.secretKey))

	if err != nil {
		return "", fmt.Errorf("cannot sign token: %w", err)
	}

	return token, nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token method")
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		return nil, fmt.Errorf("invalid token ParseWithClaims: %w", err)
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, fmt.Errorf("invalid token Claims")
	}

	return payload, nil
}
