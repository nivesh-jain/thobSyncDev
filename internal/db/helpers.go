package db

import (
	"database/sql"
	"errors"
	"fmt"
)

// GetUserByUsername retrieves a user's details by their username.
func GetUserByUsername(db *sql.DB, username string) (int, string, string, error) {
	var userID int
	var token, role string

	query := `
		SELECT u.id, u.token, ur.role
		FROM users u
		LEFT JOIN user_roles ur ON u.id = ur.user_id
		WHERE u.username = ?`

	err := db.QueryRow(query, username).Scan(&userID, &token, &role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", "", fmt.Errorf("user not found")
		}
		return 0, "", "", err
	}

	return userID, token, role, nil
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, username, token string) (int64, error) {
	query := `
		INSERT INTO users (username, token)
		VALUES (?, ?)`

	result, err := db.Exec(query, username, token)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// AssignRole assigns a role to a user.
func AssignRole(db *sql.DB, userID int, role string) error {
	query := `
		INSERT INTO user_roles (user_id, role)
		VALUES (?, ?)`

	_, err := db.Exec(query, userID, role)
	if err != nil {
		return err
	}

	return nil
}

// ListUsers retrieves all users and their roles.
func ListUsers(db *sql.DB) ([]map[string]string, error) {
	query := `
		SELECT u.username, ur.role
		FROM users u
		LEFT JOIN user_roles ur ON u.id = ur.user_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []map[string]string
	for rows.Next() {
		var username, role string
		if err := rows.Scan(&username, &role); err != nil {
			return nil, err
		}

		users = append(users, map[string]string{
			"username": username,
			"role":     role,
		})
	}

	return users, nil
}

// ListRoles retrieves all available roles.
func ListRoles(db *sql.DB) ([]string, error) {
	query := `
		SELECT name
		FROM roles`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// ValidateUserRole checks if a user has a specific role.
func ValidateUserRole(db *sql.DB, userID int, role string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM user_roles
		WHERE user_id = ? AND role = ?`

	var count int
	err := db.QueryRow(query, userID, role).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
