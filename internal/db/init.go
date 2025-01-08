package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    token TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS roles (
    name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id INTEGER NOT NULL,
    role TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role) REFERENCES roles (name)
);
`

// InitDB initializes the SQLite database with the schema.
func InitDB(dbPath string) (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Println("Creating database file...")
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to execute schema: %w", err)
	}

	fmt.Println("Database initialized successfully.")
	return db, nil
}
