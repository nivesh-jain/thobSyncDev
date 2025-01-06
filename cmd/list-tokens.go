package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var listTokensCmd = &cobra.Command{
	Use:   "list-tokens",
	Short: "List all generated tokens and their associated roles (Admin only)",
	Run: func(cmd *cobra.Command, args []string) {
		adminUsername, _ := cmd.Flags().GetString("admin")
		adminPassword, _ := cmd.Flags().GetString("password")

		if adminUsername == "" || adminPassword == "" {
			log.Fatalln("Admin username and password are required.")
		}

		if adminUsername != "admin" || adminPassword != "admin1234" {
			log.Fatalln("Access Denied: Only Admin can list tokens.")
		}

		// Read and display tokens
		tokens, err := listAllTokens()
		if err != nil {
			log.Fatalf("Failed to list tokens: %v\n", err)
		}

		fmt.Println("Generated Tokens and Associated Roles:")
		for token, role := range tokens {
			fmt.Printf("- Token: %s | Role: %s\n", token, role)
		}
	},
}

func listAllTokens() (map[string]string, error) {
	file, err := os.Open("tokens.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open token file: %v", err)
	}
	defer file.Close()

	var data struct {
		Tokens map[string]string `json:"tokens"`
	}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode token file: %v", err)
	}

	return data.Tokens, nil
}

func init() {
	listTokensCmd.Flags().StringP("admin", "a", "", "Admin username")
	listTokensCmd.Flags().StringP("password", "p", "", "Admin password")
	rootCmd.AddCommand(listTokensCmd)
}
