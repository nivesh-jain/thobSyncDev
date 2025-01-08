package cmd

import (
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"

	"github.com/spf13/cobra"
)

var username string
var token string
var dbPath string

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the database connection
		databaseConnection, err := db.InitDB(dbPath)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		defer databaseConnection.Close()

		// Authenticate the user
		valid, err := auth.AuthenticateUser(databaseConnection, username, token)
		if err != nil {
			log.Fatalf("Authentication failed: %v", err)
		}

		if valid {
			fmt.Printf("Authentication successful for user: %s\n", username)
		} else {
			fmt.Println("Invalid username or token.")
		}
	},
}

func init() {
	authCmd.Flags().StringVarP(&dbPath, "db-path", "d", "minio_cli.db", "Path to SQLite database")
	authCmd.Flags().StringVarP(&username, "username", "u", "", "Username to authenticate")
	authCmd.Flags().StringVarP(&token, "token", "t", "", "Token for authentication")
	rootCmd.AddCommand(authCmd)
}
