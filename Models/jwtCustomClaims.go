package models

import (
	"github.com/dgrijalva/jwt-go"

)

type JwtCustomClaims struct {
	Email  string `json:"email"`
	jwt.StandardClaims
}