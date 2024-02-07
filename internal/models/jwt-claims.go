package models

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	jwt.StandardClaims
}
