package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("hahahahapasswordd1321")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
