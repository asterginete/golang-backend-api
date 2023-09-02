package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SECRET_KEY = "secret_key" // Change this to a more secure key
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires after 72 hours

	return token.SignedString([]byte(SECRET_KEY))
}
