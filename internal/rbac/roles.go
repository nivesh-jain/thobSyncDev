package rbac

import "fmt"

var roles = map[string][]string{
	"Admin":  {"create-bucket", "delete-bucket", "upload-file", "download-file", "delete-file", "list-files"},
	"Editor": {"upload-file", "download-file", "delete-file", "list-files"},
	"Viewer": {"list-files"},
}

// GetPermissions retrieves permissions for a given role
func GetPermissions(role string) ([]string, error) {
	permissions, exists := roles[role]
	if !exists {
		return nil, fmt.Errorf("role '%s' does not exist", role)
	}
	return permissions, nil
}
