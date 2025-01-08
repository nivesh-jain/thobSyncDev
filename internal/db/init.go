package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const authSchema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    token TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL CHECK(role IN ('Admin', 'Editor', 'Viewer'))
);
`

const logSchema = `
CREATE TABLE IF NOT EXISTS logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER NOT NULL,
    operation TEXT NOT NULL,
    file_name TEXT,
    bucket_name TEXT,
    status TEXT CHECK(status IN ('Success', 'Failure')),
    FOREIGN KEY(user_id) REFERENCES users(id)
);
`

// InitAuthDB initializes the Authentication SQLite File.
func InitAuthDB(dbPath string) (*sql.DB, error) {
	return initDB(dbPath, authSchema)
}

// InitLogDB initializes the Logging SQLite File.
func InitLogDB(dbPath string) (*sql.DB, error) {
	return initDB(dbPath, logSchema)
}

func initDB(dbPath string, schema string) (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Printf("Creating database file: %s...\n", dbPath)
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
		return nil, fmt.Errorf("failed to apply schema: %w", err)
	}

	return db, nil
}
