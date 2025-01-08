package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var unlockFileCmd = &cobra.Command{
	Use:   "unlock-file",
	Short: "Unlock a file to allow modifications",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor")

		// Simulate file unlock
		fmt.Printf("Unlocking file '%s' in bucket '%s'...\n", filePath, bucketName)

		// Logic to unlock file (metadata update or specific mechanism)
		client := db.GetMinioClient() // Ensure GetMinioClient is implemented
		metadata := map[string]string{"locked": "false"}
		_, err := client.CopyObject(context.Background(), minio.CopyDestOptions{
			Bucket: bucketName,
			Object: filePath,
		}, minio.CopySrcOptions{
			Bucket:       bucketName,
			Object:       filePath,
			UserMetadata: metadata,
		})
		if err != nil {
			fmt.Printf("Failed to unlock file: %v\n", err)
			db.LogOperation(username, "Unlock", filePath, bucketName, "Failure")
			return
		}

		fmt.Println("File unlocked successfully!")

		// Log the operation
		db.LogOperation(username, "Unlock", filePath, bucketName, "Success")
	},
}

func init() {
	unlockFileCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
	unlockFileCmd.Flags().StringVar(&filePath, "file", "", "Path to the file to unlock")
	rootCmd.AddCommand(unlockFileCmd)
}
