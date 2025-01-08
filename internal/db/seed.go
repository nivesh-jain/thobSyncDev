package db

import "database/sql"

// SeedRoles inserts default roles into the roles table.
func SeedRoles(db *sql.DB) error {
	roles := []string{"Admin", "Editor", "Viewer"}

	stmt, err := db.Prepare("INSERT OR IGNORE INTO roles (name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, role := range roles {
		_, err = stmt.Exec(role)
		if err != nil {
			return err
		}
	}

	return nil
}
