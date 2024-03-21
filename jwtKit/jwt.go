package jwtKit

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	key []byte
}

func NewJWT(key []byte) *JWT {
	return &JWT{key: key}
}

func (j *JWT) NewWithClaims(signingMethod jwt.SigningMethod, claims jwt.Claims, opts ...jwt.TokenOption) (string, error) {
	return jwt.NewWithClaims(signingMethod, claims, opts...).SignedString(j.key)
}

func (j *JWT) Parse(tokenString string, keyFunc jwt.Keyfunc, opts ...jwt.ParserOption) (jwt.MapClaims, error) {
	token, err := jwt.NewParser(opts...).Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, nil
}
func (j *JWT) ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc, opts ...jwt.ParserOption) (*jwt.Token, error) {
	token, err := jwt.NewParser(opts...).ParseWithClaims(tokenString, claims, keyFunc)
	return token, err
}
