package auth

import (
	"database/sql"
	"errors"
	"math/rand"
	"time"
)

// GenerateToken generates a simple random token (for example purposes).
func GenerateToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

// AuthenticateUser authenticates a user by username and token.
func AuthenticateUser(db *sql.DB, username, token string) (bool, error) {
	var storedToken string
	err := db.QueryRow("SELECT token FROM users WHERE username = ?", username).Scan(&storedToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("user not found")
		}
		return false, err
	}

	if storedToken != token {
		return false, errors.New("invalid token")
	}

	return true, nil
}

// AssignRole assigns a role to a user.
func AssignRole(db *sql.DB, username, role string) error {
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	_, err = db.Exec("INSERT INTO user_roles (user_id, role) VALUES (?, ?)", userID, role)
	if err != nil {
		return err
	}

	return nil
}
