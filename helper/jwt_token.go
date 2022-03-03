package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenJwt struct{}

type TokenJWT interface {
	Generate(id uint) (string, error)
	Decode(string) (interface{}, error)
}

type myCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

var mySigningKey []byte = []byte(os.Getenv("SECRET_KEY"))

func NewTokenJWT() TokenJWT {
	return tokenJwt{}
}

func (tn tokenJwt) Generate(id uint) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := myCustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "test",
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (token tokenJwt) Decode(tokenString string) (interface{}, error) {
	claims := &myCustomClaims{}
	decode, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !decode.Valid {
		return nil, errors.New("token tidak valid")
	}
	return claims.ID, nil
}
