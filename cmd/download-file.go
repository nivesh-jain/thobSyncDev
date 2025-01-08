package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var downloadBucketName string
var downloadFilePath string

var downloadFileCmd = &cobra.Command{
	Use:   "download-file",
	Short: "Download a file from the specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Log the operation
		fmt.Printf("Downloading file '%s' from bucket '%s'...\n", downloadFilePath, downloadBucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// Download the file
		err := client.FGetObject(
			context.Background(),
			downloadBucketName,
			downloadFilePath, // Object name
			downloadFilePath, // Destination path
			minio.GetObjectOptions{},
		)
		if err != nil {
			fmt.Printf("Failed to download file: %v\n", err)
			db.LogOperation(username, "Download", downloadFilePath, downloadBucketName, "Failure")
			return
		}

		fmt.Println("File downloaded successfully!")
		db.LogOperation(username, "Download", downloadFilePath, downloadBucketName, "Success")
	},
}

func init() {
	downloadFileCmd.Flags().StringVar(&downloadBucketName, "bucket", "", "Name of the bucket")
	downloadFileCmd.Flags().StringVar(&downloadFilePath, "file", "", "Path to the file to download")
	rootCmd.AddCommand(downloadFileCmd)
}
