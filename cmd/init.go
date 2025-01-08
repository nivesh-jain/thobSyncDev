package cmd

import (
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"

	"github.com/spf13/cobra"
)

var authDBPath string
var logDBPath string
var adminUsername string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application as an Admin",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize Authentication DB
		authDB, err := db.InitAuthDB(authDBPath)
		if err != nil {
			log.Fatalf("Failed to initialize Authentication DB: %v", err)
		}
		defer authDB.Close()

		// Initialize Logging DB
		logDB, err := db.InitLogDB(logDBPath)
		if err != nil {
			log.Fatalf("Failed to initialize Logging DB: %v", err)
		}
		defer logDB.Close()

		// Seed Admin User
		token := auth.GenerateToken()
		_, err = authDB.Exec("INSERT INTO users (username, token, role) VALUES (?, ?, ?)", adminUsername, token, "Admin")
		if err != nil {
			log.Fatalf("Failed to create Admin user: %v", err)
		}

		fmt.Printf("Admin user '%s' created successfully.\nToken: %s\n", adminUsername, token)

		// TODO: Upload SQLite files to S3 bucket
	},
}

func init() {
	initCmd.Flags().StringVar(&authDBPath, "auth-db", "auth.db", "Path to Authentication DB")
	initCmd.Flags().StringVar(&logDBPath, "log-db", "log.db", "Path to Logging DB")
	initCmd.Flags().StringVar(&adminUsername, "admin", "admin", "Admin username")
	rootCmd.AddCommand(initCmd)
}
