package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(iss, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": iss,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func main() {
	fmt.Println(CreateToken("keystone", "R0ckso1id"))
}
