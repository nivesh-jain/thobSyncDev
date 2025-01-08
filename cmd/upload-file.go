package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/config"
	"github.com/spf13/cobra"
)

var bucketName string
var filePath string

var uploadFileCmd = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload a file to the specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username, role := config.GetCurrentUser()
		if role != "Admin" && role != "Editor" {
			log.Fatalf("Permission denied: %s role cannot upload files.", role)
		}

		// Simulate file upload
		fmt.Printf("Uploading file '%s' to bucket '%s'...\n", filePath, bucketName)
		// Add logic to interact with S3 bucket
		fmt.Println("File uploaded successfully!")

		// Log the operation
		logDB, err := sql.Open("sqlite3", logDBPath)
		if err != nil {
			log.Fatalf("Failed to open Logging DB: %v", err)
		}
		defer logDB.Close()

		_, err = logDB.Exec("INSERT INTO logs (user_id, operation, file_name, bucket_name, status) VALUES (?, ?, ?, ?, ?)",
			username, "Upload", filePath, bucketName, "Success")
		if err != nil {
			log.Fatalf("Failed to log operation: %v", err)
		}
	},
}

func init() {
	uploadFileCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
	uploadFileCmd.Flags().StringVar(&filePath, "file", "", "Path to the file")
	rootCmd.AddCommand(uploadFileCmd)
}
