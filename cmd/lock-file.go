package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var lockFileCmd = &cobra.Command{
	Use:   "lock-file",
	Short: "Lock a file to prevent modifications",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Simulate file lock
		fmt.Printf("Locking file '%s' in bucket '%s'...\n", filePath, bucketName)

		// Logic to lock file (metadata update or specific mechanism)
		client := db.GetMinioClient() // Ensure GetMinioClient is implemented
		metadata := map[string]string{"locked": "true"}
		_, err := client.CopyObject(context.Background(), minio.CopyDestOptions{
			Bucket: bucketName,
			Object: filePath,
		}, minio.CopySrcOptions{
			Bucket:       bucketName,
			Object:       filePath,
			UserMetadata: metadata,
		})
		if err != nil {
			fmt.Printf("Failed to lock file: %v\n", err)
			db.LogOperation(username, "Lock", filePath, bucketName, "Failure")
			return
		}

		fmt.Println("File locked successfully!")

		// Log the operation
		db.LogOperation(username, "Lock", filePath, bucketName, "Success")
	},
}

func init() {
	lockFileCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
	lockFileCmd.Flags().StringVar(&filePath, "file", "", "Path to the file to lock")
	rootCmd.AddCommand(lockFileCmd)
}
