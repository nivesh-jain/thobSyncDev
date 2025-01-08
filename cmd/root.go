package cmd

import (
	"database/sql"
	"log"

	"github.com/spf13/cobra"
)

var authDBPath string
var databaseConnection *sql.DB

var rootCmd = &cobra.Command{
	Use:   "thobSyncDev",
	Short: "CLI tool for managing files with authentication and logging",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize the Authentication DB for all commands
		var err error
		databaseConnection, err = sql.Open("sqlite3", authDBPath)
		if err != nil {
			log.Fatalf("Failed to open Authentication DB: %v", err)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Close the database connection
		if databaseConnection != nil {
			databaseConnection.Close()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func init() {
	// Add global flags
	rootCmd.PersistentFlags().StringVar(&authDBPath, "auth-db", "auth.db", "Path to the Authentication DB")
}
