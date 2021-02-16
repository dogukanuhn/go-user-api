package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"user-basic/models"
)





func Authenticate( userMail string) (string,error)  {

	// Set custom claims
	claims := &models.JwtCustomClaims{
		userMail,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "",err
	}

	return t,err
}