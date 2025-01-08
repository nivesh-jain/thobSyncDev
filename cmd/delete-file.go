package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var deleteFileCmd = &cobra.Command{
	Use:   "delete-file",
	Short: "Delete a file from the specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Simulate file deletion
		fmt.Printf("Deleting file '%s' from bucket '%s'...\n", filePath, bucketName)

		// Logic to delete file
		client := db.GetMinioClient() // Ensure GetMinioClient is implemented
		err := client.RemoveObject(context.Background(), bucketName, filePath, minio.RemoveObjectOptions{})
		if err != nil {
			fmt.Printf("Failed to delete file: %v\n", err)
			db.LogOperation(username, "Delete", filePath, bucketName, "Failure")
			return
		}

		fmt.Println("File deleted successfully!")

		// Log the operation
		db.LogOperation(username, "Delete", filePath, bucketName, "Success")
	},
}

func init() {
	deleteFileCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
	deleteFileCmd.Flags().StringVar(&filePath, "file", "", "Path to the file to delete")
	rootCmd.AddCommand(deleteFileCmd)
}
