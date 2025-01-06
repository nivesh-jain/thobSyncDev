package cmd

import (
	"log"
	"path/filepath"

	"github.com/nivesh-jain/thobSyncDev.git/internal/minio"
	"github.com/nivesh-jain/thobSyncDev.git/internal/rbac"
	"github.com/spf13/cobra"
)

var uploadFileCmd = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload a file to a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		role, _ := cmd.Flags().GetString("role")

		bucketName, _ := cmd.Flags().GetString("bucket")
		filePath, _ := cmd.Flags().GetString("file")

		if role == "" {
			log.Fatalln("Role is required. Use the --role flag to specify a role (e.g., Admin, Editor, Viewer).")
		}

		if bucketName == "" || filePath == "" {
			log.Fatalln("Bucket name and file path are required.")
		}
		// Check permission
		if err := rbac.CheckPermission(role, "upload-file"); err != nil {
			log.Fatalf("Access Denied: %v\n", err)
		}

		client := minio.NewClient()
		objectName := filepath.Base(filePath)

		minio.UploadFile(client, bucketName, objectName, filePath)
	},
}

func init() {
	uploadFileCmd.Flags().StringP("role", "r", "", "Role of the user (Admin, Editor, Viewer)")
	uploadFileCmd.Flags().StringP("bucket", "b", "", "Name of the bucket")
	uploadFileCmd.Flags().StringP("file", "f", "", "Path to the file to upload")
	rootCmd.AddCommand(uploadFileCmd)
}
