package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Expire time.Duration
	Secret string
)

type JWTClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

func GenToken(claims *JWTClaims) (string, error) {
	c := JWTClaims{
		Account: claims.Account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(Expire).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(Secret))
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return []byte(Secret), nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is invalid")
}
