package jwtKit

import (
	"fmt"
	"log"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestNewJWT(t *testing.T) {
	var signKey = []byte("rabbit.RM")
	claims := &jwt.RegisteredClaims{
		Issuer: "RM",
		ID:     "1",
	}
	text, err := NewWithClaims(signKey, jwt.SigningMethodHS256, claims)
	fmt.Println(text, err)
	v, err := ParseWithClaims(text, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v.Issuer, v.ID)
}
