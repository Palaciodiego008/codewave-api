package services

import (
	"codewave/config"
	"codewave/models"
	"codewave/utils"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

func CreateUser(user *models.User) error {
	// Hash the password
	hashedPassword, err := utils.GenerateHash(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword
	if config.DB == nil {
		return errors.New("database connection not initialized")
	}

	err = config.DB.Create(user).Error
	return err
}

func GetUser(id string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func AuthenticateUser(email string, password string) (string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Print stored hashed password
	fmt.Println("Stored hashed password:", user.Password)

	// Compare the stored hashed password with the password provided
	err := utils.ComparePasswordHash(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	// Ensure jwtKey is defined
	if jwtKey == nil {
		return "", errors.New("JWT key not defined")
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
