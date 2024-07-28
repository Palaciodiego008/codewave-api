package utils

import (
	"codewave/models"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

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

// Function to extract JSON from a string
func ExtractJSON(input string) (models.GeminiResponse, error) {
	var result models.GeminiResponse

	// Regular expression to find JSON content
	re := regexp.MustCompile(`(?s)\{.*\}`)
	jsonStr := re.FindString(input)

	if jsonStr == "" {
		return result, fmt.Errorf("no JSON found in input")
	}

	// Clean up the JSON string
	jsonStr = strings.TrimSpace(jsonStr)
	jsonStr = strings.TrimPrefix(jsonStr, "{[```json")
	jsonStr = strings.TrimSuffix(jsonStr, "```]")
	jsonStr = strings.TrimSuffix(jsonStr, "model}")

	// Remove any extraneous characters
	jsonStr = strings.Trim(jsonStr, "{} ")
	jsonStr = strings.TrimPrefix(jsonStr, "{")
	jsonStr = strings.TrimSuffix(jsonStr, "}")
	jsonStr = strings.TrimSuffix(jsonStr, "```]")

	fmt.Println("Extracted JSON: ", jsonStr)

	// Unmarshal the JSON string into the struct
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return result, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return result, nil
}

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
