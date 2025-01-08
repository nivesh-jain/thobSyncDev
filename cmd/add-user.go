package cmd

import (
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/spf13/cobra"
)

var newUsername string
var userRole string

var addUserCmd = &cobra.Command{
	Use:   "add-user",
	Short: "Add a new user to the application",
	Run: func(cmd *cobra.Command, args []string) {
		// Use the global `databaseConnection` initialized in root
		token := auth.GenerateToken()

		_, err := databaseConnection.Exec("INSERT INTO users (username, token, role) VALUES (?, ?, ?)", newUsername, token, userRole)
		if err != nil {
			log.Fatalf("Failed to add user: %v", err)
		}

		fmt.Printf("User '%s' added successfully with role '%s'.\nToken: %s\n", newUsername, userRole, token)

		// TODO: Upload the updated Authentication DB to S3 bucket
	},
}

func init() {
	addUserCmd.Flags().StringVar(&newUsername, "username", "", "New user's username")
	addUserCmd.Flags().StringVar(&userRole, "role", "Editor", "Role for the new user (Admin, Editor, Viewer)")
	rootCmd.AddCommand(addUserCmd)
}
