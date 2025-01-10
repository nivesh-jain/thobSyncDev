package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var infoBucketName string
var infoFilePath string

var getInfoCmd = &cobra.Command{
	Use:   "get-info",
	Short: "Fetch metadata, retention info, and legal hold status for a file",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor", "Viewer")

		// Log the operation
		fmt.Printf("Fetching information for file '%s' in bucket '%s'...\n", infoFilePath, infoBucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// Fetch object info
		stat, err := client.StatObject(
			context.Background(),
			infoBucketName,
			infoFilePath,
			minio.StatObjectOptions{},
		)
		if err != nil {
			fmt.Printf("Failed to fetch file information: %v\n", err)
			db.LogOperation(username, "GetInfo", infoFilePath, infoBucketName, "Failure")
			return
		}

		// Display general metadata
		fmt.Println("File Information:")
		fmt.Printf("  Name: %s\n", stat.Key)
		fmt.Printf("  Size: %d bytes\n", stat.Size)
		fmt.Printf("  Content-Type: %s\n", stat.ContentType)
		fmt.Printf("  Last Modified: %v\n", stat.LastModified)

		// Display custom metadata
		fmt.Println("Custom Metadata:")
		for key, value := range stat.UserMetadata {
			fmt.Printf("  %s: %s\n", key, value)
		}

		// Fetch retention information
		retention, retainUntilDate, err := client.GetObjectRetention(
			context.Background(),
			infoBucketName,
			infoFilePath,
			"",
		)
		if err != nil {
			fmt.Printf("Failed to get retention info: %v\n", err)
		} else {
			fmt.Printf("Retention Info: %s %v", retention, retainUntilDate)
		}

		// Fetch legal hold information
		legalHold, err := client.GetObjectLegalHold(
			context.Background(),
			infoBucketName,
			infoFilePath,
			minio.GetObjectLegalHoldOptions{
				VersionID: "",
			},
		)
		if err != nil {
			fmt.Printf("Failed to get legal hold info: %v\n", err)
		} else {
			fmt.Printf("Legal Hold: %v\n", legalHold)
		}

		db.LogOperation(username, "GetInfo", infoFilePath, infoBucketName, "Success")
	},
}

func init() {
	getInfoCmd.Flags().StringVar(&infoBucketName, "bucket", "", "Name of the bucket")
	getInfoCmd.Flags().StringVar(&infoFilePath, "file", "", "Path to the file to fetch information")
	rootCmd.AddCommand(getInfoCmd)
}
