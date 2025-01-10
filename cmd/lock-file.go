package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var lockBucketName string
var lockFilePath string

var lockFileCmd = &cobra.Command{
	Use:   "lock-file",
	Short: "Lock a file in a bucket to prevent modifications",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Log the operation
		fmt.Printf("Locking file '%s' in bucket '%s'...\n", lockFilePath, lockBucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// Define custom metadata
		userMetadata := map[string]string{
			"locked-by": username,
		}

		// Define retention options
		mode := minio.Governance
		retainUntilDate := time.Now().Add(24 * time.Hour) // Lock for 24 hours

		// Apply retention, legal hold, and metadata updates using CopyObject
		_, err := client.CopyObject(
			context.Background(),
			minio.CopyDestOptions{
				Bucket:          lockBucketName,
				Object:          lockFilePath,
				ReplaceMetadata: true,                   // Force metadata replacement
				UserMetadata:    userMetadata,           // Update metadata
				LegalHold:       minio.LegalHoldEnabled, // Enable legal hold
				Mode:            mode,                   // Set governance mode
				RetainUntilDate: retainUntilDate,        // Set retain-until date
			},
			minio.CopySrcOptions{
				Bucket: lockBucketName,
				Object: lockFilePath,
			},
		)
		if err != nil {
			fmt.Printf("Failed to apply lock and update metadata: %v\n", err)
			db.LogOperation(username, "Lock", lockFilePath, lockBucketName, "Failure")
			return
		}

		// Log success
		fmt.Println("File locked successfully with retention and updated metadata!")
		db.LogOperation(username, "Lock", lockFilePath, lockBucketName, "Success")
	},
}

func init() {
	lockFileCmd.Flags().StringVar(&lockBucketName, "bucket", "", "Name of the bucket")
	lockFileCmd.Flags().StringVar(&lockFilePath, "file", "", "Path to the file to lock")
	rootCmd.AddCommand(lockFileCmd)
}
