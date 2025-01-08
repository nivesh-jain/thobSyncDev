package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var lockBucketName string
var lockFilePath string

var lockFileCmd = &cobra.Command{
	Use:   "lock-file",
	Short: "Lock a file to prevent modifications",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Log the operation
		fmt.Printf("Locking file '%s' in bucket '%s'...\n", lockFilePath, lockBucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// Copy the object with updated metadata
		_, err := client.CopyObject(
			context.Background(),
			minio.CopyDestOptions{
				Bucket:          lockBucketName,
				Object:          lockFilePath,
				ReplaceMetadata: true, // Ensure metadata is replaced
				UserMetadata:    map[string]string{"locked": "true"},
			},
			minio.CopySrcOptions{
				Bucket: lockBucketName,
				Object: lockFilePath,
			},
		)
		if err != nil {
			fmt.Printf("Failed to lock file: %v\n", err)
			db.LogOperation(username, "Lock", lockFilePath, lockBucketName, "Failure")
			return
		}

		fmt.Println("File locked successfully!")
		db.LogOperation(username, "Lock", lockFilePath, lockBucketName, "Success")
	},
}

func init() {
	lockFileCmd.Flags().StringVar(&lockBucketName, "bucket", "", "Name of the bucket")
	lockFileCmd.Flags().StringVar(&lockFilePath, "file", "", "Path to the file to lock")
	rootCmd.AddCommand(lockFileCmd)
}
