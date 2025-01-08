package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var listFilesBucketName string

var listFilesCmd = &cobra.Command{
	Use:   "list-files",
	Short: "List all files in the specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Check user's role
		username := auth.CheckUserRole("Admin", "Editor", "Viewer")

		// Log the operation
		fmt.Printf("Listing files in bucket '%s'...\n", listFilesBucketName)

		// Get MinIO client
		client := db.GetMinioClient()

		// List objects in the bucket
		objectCh := client.ListObjects(context.Background(), listFilesBucketName, minio.ListObjectsOptions{
			Recursive: true, // Recursive listing
		})

		// Iterate and display objects
		for object := range objectCh {
			if object.Err != nil {
				fmt.Printf("Error while listing object: %v\n", object.Err)
				db.LogOperation(username, "ListFiles", "", listFilesBucketName, "Failure")
				return
			}

			fmt.Printf("File: %s (Size: %d bytes)\n", object.Key, object.Size)
		}

		fmt.Println("Files listed successfully!")
		db.LogOperation(username, "ListFiles", "", listFilesBucketName, "Success")
	},
}

func init() {
	listFilesCmd.Flags().StringVar(&listFilesBucketName, "bucket", "", "Name of the bucket")
	rootCmd.AddCommand(listFilesCmd)
}
