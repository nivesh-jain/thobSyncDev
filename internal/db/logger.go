package db

import (
	"database/sql"
	"log"
)

var logDBPath = "log.db" // Path to Logging SQLite File

// LogOperation records an operation in the Logging DB.
func LogOperation(userID, operation, fileName, bucketName, status string) {
	logDB, err := sql.Open("sqlite3", logDBPath)
	if err != nil {
		log.Fatalf("Failed to open Logging DB: %v", err)
	}
	defer logDB.Close()

	_, err = logDB.Exec("INSERT INTO logs (user_id, operation, file_name, bucket_name, status) VALUES (?, ?, ?, ?, ?)",
		userID, operation, fileName, bucketName, status)
	if err != nil {
		log.Fatalf("Failed to log operation: %v", err)
	}
	log.Printf("Logged operation: %s by user: %s", operation, userID)
}
