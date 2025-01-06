package auth

import (
	"encoding/json"
	"errors"
	"os"
)

var tokenFile = "tokens.json"

// ValidateToken checks if a token is valid and returns the associated role
func ValidateToken(token string) (string, error) {
	file, err := os.Open(tokenFile)
	if err != nil {
		return "", errors.New("failed to open token file")
	}
	defer file.Close()

	var data struct {
		Tokens map[string]string `json:"tokens"`
	}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return "", errors.New("failed to decode token file")
	}

	role, exists := data.Tokens[token]
	if !exists {
		return "", errors.New("invalid token")
	}

	return role, nil
}
