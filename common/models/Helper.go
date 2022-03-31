package models

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
