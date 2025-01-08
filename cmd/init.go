package cmd

import (
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"

	"github.com/spf13/cobra"
)

// initCmd defines the "init" command for initializing the CLI with a new user.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI with a new user",
	Run: func(cmd *cobra.Command, args []string) {
		// Generate a random token for the user
		token := auth.GenerateToken()

		// Initialize the database
		databaseConnection, err := db.InitDB(dbPath)
		if err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}
		defer databaseConnection.Close()

		// Seed roles if not already present
		err = db.SeedRoles(databaseConnection)
		if err != nil {
			log.Fatalf("Failed to seed roles: %v", err)
		}

		// Check if the username already exists
		_, _, _, err = db.GetUserByUsername(databaseConnection, username)
		if err == nil {
			log.Fatalf("User '%s' already exists.", username)
		}

		// Insert the new user into the database
		userID, err := db.CreateUser(databaseConnection, username, token)
		if err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}

		// Assign "Admin" role to the first user created
		err = db.AssignRole(databaseConnection, int(userID), "Admin")
		if err != nil {
			log.Fatalf("Failed to assign 'Admin' role to user: %v", err)
		}

		fmt.Printf("User '%s' created successfully with 'Admin' role.\n", username)
		fmt.Printf("Token: %s\n", token)
	},
}

func init() {
	// Define flags for the init command
	initCmd.Flags().StringVarP(&dbPath, "db-path", "d", "minio_cli.db", "Path to SQLite database")
	initCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the new user")
	rootCmd.AddCommand(initCmd)
}
