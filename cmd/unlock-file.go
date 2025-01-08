package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var unlockBucketName string
var unlockFilePath string

var unlockFileCmd = &cobra.Command{
	Use:   "unlock-file",
	Short: "Unlock a file to allow modifications",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Log the operation
		fmt.Printf("Unlocking file '%s' in bucket '%s'...\n", unlockFilePath, unlockBucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// Copy the object with updated metadata
		_, err := client.CopyObject(
			context.Background(),
			minio.CopyDestOptions{
				Bucket:          unlockBucketName,
				Object:          unlockFilePath,
				ReplaceMetadata: true, // Ensure metadata is replaced
				UserMetadata:    map[string]string{"locked": "false"},
			},
			minio.CopySrcOptions{
				Bucket: unlockBucketName,
				Object: unlockFilePath,
			},
		)
		if err != nil {
			fmt.Printf("Failed to unlock file: %v\n", err)
			db.LogOperation(username, "Unlock", unlockFilePath, unlockBucketName, "Failure")
			return
		}

		fmt.Println("File unlocked successfully!")
		db.LogOperation(username, "Unlock", unlockFilePath, unlockBucketName, "Success")
	},
}

func init() {
	unlockFileCmd.Flags().StringVar(&unlockBucketName, "bucket", "", "Name of the bucket")
	unlockFileCmd.Flags().StringVar(&unlockFilePath, "file", "", "Path to the file to unlock")
	rootCmd.AddCommand(unlockFileCmd)
}
