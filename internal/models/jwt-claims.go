package models

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	Names []string `json:"names"`
	jwt.StandardClaims
}
