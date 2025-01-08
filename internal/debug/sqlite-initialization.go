package main

import (
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
)

func main() {
	authDB, err := db.InitAuthDB("auth.db")
	if err != nil {
		log.Fatalf("Failed to initialize Authentication DB: %v", err)
	}
	defer authDB.Close()

	logDB, err := db.InitLogDB("log.db")
	if err != nil {
		log.Fatalf("Failed to initialize Logging DB: %v", err)
	}
	defer logDB.Close()

	log.Println("Databases initialized successfully!")
}
