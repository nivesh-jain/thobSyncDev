package cmd

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var downloadBucketName string
var downloadFilePath string
var destinationPath string

var downloadFileCmd = &cobra.Command{
	Use:   "download-file",
	Short: "Download a file from a specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor", "Viewer")

		// Get MinIO client
		client := db.GetMinioClient()

		// Fetch object metadata
		stat, err := client.StatObject(
			context.Background(),
			downloadBucketName,
			downloadFilePath,
			minio.StatObjectOptions{},
		)
		if err != nil {
			log.Fatalf("Failed to fetch file information: %v\n", err)
		}

		// Check retention info
		retention, retainUntilDate, err := client.GetObjectRetention(
			context.Background(),
			downloadBucketName,
			downloadFilePath,
			"",
		)
		if err == nil && retention != nil {
			fmt.Printf("File is locked under %s mode until %v.\n", retention, retainUntilDate)
			lockedBy := stat.UserMetadata["Locked-By"]
			if lockedBy != "" {
				fmt.Printf("File was locked by: %s\n", lockedBy)
			}
			log.Println("Download operation aborted due to file lock.")
			return
		}

		// Construct destination file path
		fileName := filepath.Base(downloadFilePath) // Extract only the file name
		dest := filepath.Join(destinationPath, fileName)

		// Proceed with download
		err = client.FGetObject(
			context.Background(),
			downloadBucketName,
			downloadFilePath,
			dest,
			minio.GetObjectOptions{},
		)
		if err != nil {
			log.Fatalf("Failed to download file: %v\n", err)
		}

		// Log success
		fmt.Printf("File '%s' downloaded successfully to '%s'.\n", downloadFilePath, dest)
		db.LogOperation(username, "Download", downloadFilePath, downloadBucketName, "Success")
	},
}

func init() {
	downloadFileCmd.Flags().StringVar(&downloadBucketName, "bucket", "", "Name of the bucket")
	downloadFileCmd.Flags().StringVar(&downloadFilePath, "file", "", "Path to the file to download")
	downloadFileCmd.Flags().StringVar(&destinationPath, "dest", ".", "Destination directory path")
	rootCmd.AddCommand(downloadFileCmd)
}
