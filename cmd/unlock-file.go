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

var unlockBucketName string
var unlockFilePath string

var unlockFileCmd = &cobra.Command{
	Use:   "unlock-file",
	Short: "Unlock a file in a bucket to allow modifications",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin")

		// Log the operation
		fmt.Printf("Unlocking file '%s' in bucket '%s'...\n", unlockFilePath, unlockBucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// Update retention settings to remove lock
		_, err := client.CopyObject(
			context.Background(),
			minio.CopyDestOptions{
				Bucket:          unlockBucketName,
				Object:          unlockFilePath,
				ReplaceMetadata: true,                // Force metadata replacement
				UserMetadata:    map[string]string{}, // Provide empty metadata to avoid conflicts
				Mode:            "",                  // Remove retention policy
				RetainUntilDate: time.Time{},         // Clear retention date
			},
			minio.CopySrcOptions{
				Bucket: unlockBucketName,
				Object: unlockFilePath,
			},
		)
		if err != nil {
			fmt.Printf("Failed to remove retention policy: %v\n", err)
			db.LogOperation(username, "Unlock", unlockFilePath, unlockBucketName, "Failure")
			return
		}

		// Disable legal hold
		_, err = client.CopyObject(
			context.Background(),
			minio.CopyDestOptions{
				Bucket:          unlockBucketName,
				Object:          unlockFilePath,
				LegalHold:       minio.LegalHoldDisabled, // Disable legal hold
				ReplaceMetadata: true,                    // Ensure metadata is updated
			},
			minio.CopySrcOptions{
				Bucket: unlockBucketName,
				Object: unlockFilePath,
			},
		)
		if err != nil {
			fmt.Printf("Failed to remove legal hold: %v\n", err)
			db.LogOperation(username, "Unlock", unlockFilePath, unlockBucketName, "Failure")
			return
		}

		// Log success
		fmt.Println("File unlocked successfully!")
		db.LogOperation(username, "Unlock", unlockFilePath, unlockBucketName, "Success")
	},
}

func init() {
	unlockFileCmd.Flags().StringVar(&unlockBucketName, "bucket", "", "Name of the bucket")
	unlockFileCmd.Flags().StringVar(&unlockFilePath, "file", "", "Path to the file to unlock")
	rootCmd.AddCommand(unlockFileCmd)
}
