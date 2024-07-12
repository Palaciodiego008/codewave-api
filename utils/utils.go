package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt.
func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash compares the plain password with the hashed password.
func ComparePasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
