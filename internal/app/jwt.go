package app

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
