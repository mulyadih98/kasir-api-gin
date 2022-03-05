package helper

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenJwt struct{}

type TokenJWT interface {
	Generate(id uint) (string, error)
	GenerateRefresh(id uint) (string, error)
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
	expired, _ := strconv.ParseInt(os.Getenv("TOKEN_EXPIRATION"), 10, 32)
	expirationTime := time.Now().Add(time.Duration(expired) * time.Minute)
	claims := myCustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tn tokenJwt) GenerateRefresh(id uint) (string, error) {
	claims := myCustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			Issuer: os.Getenv("APP_NAME"),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	RefreshTokenString, err := jwtToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return RefreshTokenString, nil
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
