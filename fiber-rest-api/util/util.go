package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
)

func GetEmailFromToken(authorization string) (string, error) {
	tokenStr := authorization[7:len(authorization)]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, &jwt.StandardClaims{})
	if err != nil {
		log.Println(err)
		return "", err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("cannot parse claims")
	}
	return claims.Subject, nil
}
