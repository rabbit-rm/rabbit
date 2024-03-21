package jwtKit

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestNewJWT(t *testing.T) {
	var signKey = []byte("rabbit.RM")
	claims := &jwt.RegisteredClaims{
		Issuer: "RM",
		ID:     "1",
	}
	j := NewJWT(signKey)
	text, err := j.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(text, err)
	token, err := j.ParseWithClaims(text, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if v, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		fmt.Println(v.Issuer, v.ID)
	}
}

func TestPathJoin(t *testing.T) {
	root := "W:\\A\\B\\C\\D\\E"
	join := path.Join(root, "..\\..\\DD")
	err := os.Chdir(join)
	if err != nil {
		log.Fatal(err)
	}
	getwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(join)
	fmt.Println(getwd)

}
