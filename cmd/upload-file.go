package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var bucketName string
var filePath string

var uploadFileCmd = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload a file to the specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Log the operation
		fmt.Printf("Uploading file '%s' to bucket '%s'...\n", filePath, bucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// Upload the file
		_, err := client.FPutObject(
			context.Background(),
			bucketName,
			filePath, // Object name (uses the file's base name)
			filePath, // Path to the file
			minio.PutObjectOptions{},
		)
		if err != nil {
			fmt.Printf("Failed to upload file: %v\n", err)
			db.LogOperation(username, "Upload", filePath, bucketName, "Failure")
			return
		}

		fmt.Println("File uploaded successfully!")
		db.LogOperation(username, "Upload", filePath, bucketName, "Success")
	},
}

func init() {
	uploadFileCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
	uploadFileCmd.Flags().StringVar(&filePath, "file", "", "Path to the file")
	rootCmd.AddCommand(uploadFileCmd)
}
