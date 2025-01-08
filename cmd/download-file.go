package cmd

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var downloadFileCmd = &cobra.Command{
	Use:   "download-file",
	Short: "Download a file from the specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Simulate file download
		fmt.Printf("Downloading file '%s' from bucket '%s'...\n", filePath, bucketName)

		// Logic to download file
		client := db.GetMinioClient() // Ensure GetMinioClient is implemented to initialize MinIO client
		err := client.FGetObject(bucketName, filePath, filePath, minio.GetObjectOptions{})
		if err != nil {
			fmt.Printf("Failed to download file: %v\n", err)
			db.LogOperation(username, "Download", filePath, bucketName, "Failure")
			return
		}

		fmt.Println("File downloaded successfully!")

		// Log the operation
		db.LogOperation(username, "Download", filePath, bucketName, "Success")
	},
}

func init() {
	downloadFileCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
	downloadFileCmd.Flags().StringVar(&filePath, "file", "", "Path to the file to download")
	rootCmd.AddCommand(downloadFileCmd)
}
