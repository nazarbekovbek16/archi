package model

import "github.com/golang-jwt/jwt"

type JWTClaim struct {
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	jwt.StandardClaims
}
