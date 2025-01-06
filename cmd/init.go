package cmd

import (
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI tool with an API token",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		if token == "" {
			log.Fatalln("API token is required.")
		}

		// Validate token and get role
		role, err := auth.ValidateToken(token)
		if err != nil {
			log.Fatalf("Failed to validate token: %v\n", err)
		}

		// Save role to config.yaml
		viper.Set("role", role)
		if err := viper.WriteConfig(); err != nil {
			log.Fatalf("Failed to save configuration: %v\n", err)
		}

		fmt.Printf("Initialization successful! Role: %s\n", role)
	},
}

func init() {
	initCmd.Flags().StringP("token", "t", "", "API token for authentication")
	rootCmd.AddCommand(initCmd)
}
