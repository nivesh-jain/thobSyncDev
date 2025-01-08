package auth

import (
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/config"
)

// CheckUserRole validates if the current user has the required role.
func CheckUserRole(requiredRoles ...string) string {
	username, role := config.GetCurrentUser()

	for _, requiredRole := range requiredRoles {
		if role == requiredRole {
			return username
		}
	}

	log.Fatalf("Permission denied: %s role cannot perform this operation.", role)
	return "" // This line will never be reached due to log.Fatalf
}
