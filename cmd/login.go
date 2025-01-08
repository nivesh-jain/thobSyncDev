package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/config"
	"github.com/spf13/cobra"
)

var userToken string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		// Open the Authentication DB
		authDB, err := sql.Open("sqlite3", authDBPath)
		if err != nil {
			log.Fatalf("Failed to open Authentication DB: %v", err)
		}
		defer authDB.Close()

		// Authenticate user
		var username, role string
		err = authDB.QueryRow("SELECT username, role FROM users WHERE token = ?", userToken).Scan(&username, &role)
		if err != nil {
			log.Fatalf("Authentication failed: %v", err)
		}

		// Cache credentials locally
		config.GenerateConfigFile(username, role, userToken)

		fmt.Printf("Login successful! Welcome, %s (Role: %s).\n", username, role)
	},
}

func init() {
	loginCmd.Flags().StringVar(&userToken, "token", "", "Authentication token")
	rootCmd.AddCommand(loginCmd)
}
