package jwtToolkit

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/rabbit-rm/rabbit/errorToolkit"
)

func NewWithClaims(secretKey interface{}, signingMethod jwt.SigningMethod, claims jwt.Claims, opts ...jwt.TokenOption) (string, error) {
	return jwt.NewWithClaims(signingMethod, claims, opts...).SignedString(secretKey)
}

func ParseWithClaims[T jwt.Claims](tokenString string, claims T, keyFunc jwt.Keyfunc, opts ...jwt.ParserOption) (value T, err error) {
	token, err := jwt.NewParser(opts...).ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid {
		err = errorToolkit.New("invalid token")
		return
	}
	var ok bool
	if value, ok = token.Claims.(T); !ok {
		err = errorToolkit.New("invalid claims type(%T)", token.Claims)
		return
	}
	return
}
