package util

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain-text password.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword verifies a password against its hash.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ... Potentially add other password utility functions, like generating random passwords, etc. ...
