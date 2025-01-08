package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {
// 	dbPath := "minio_cli.db"

// 	db, err := sql.Open("sqlite3", dbPath)
// 	if err != nil {
// 		log.Fatalf("Failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		log.Fatalf("Failed to query users: %v", err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("Users Table:")
// 	for rows.Next() {
// 		var id int
// 		var username, token string
// 		err = rows.Scan(&id, &username, &token)
// 		if err != nil {
// 			log.Fatalf("Failed to scan row: %v", err)
// 		}
// 		fmt.Printf("ID: %d, Username: %s, Token: %s\n", id, username, token)
// 	}
// }
