package rbac

import "errors"

// CheckPermission verifies if a role has the required permission
func CheckPermission(role, permission string) error {
	permissions, err := GetPermissions(role)
	if err != nil {
		return err
	}

	for _, p := range permissions {
		if p == permission {
			return nil
		}
	}
	return errors.New("permission denied")
}
