package cmd

import (
	"fmt"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var listFilesCmd = &cobra.Command{
	Use:   "list-files",
	Short: "List all files in the specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor", "Viewer")

		// Simulate file listing
		fmt.Printf("Listing files in bucket '%s'...\n", bucketName)
		// Add logic to interact with S3 bucket
		fmt.Println("Files: [example1.txt, example2.txt, example3.txt]")

		// Log the operation
		db.LogOperation(username, "ListFiles", "", bucketName, "Success")
	},
}

func init() {
	listFilesCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
	rootCmd.AddCommand(listFilesCmd)
}
