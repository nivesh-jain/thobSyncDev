package cmd

import (
	"log"
	"path/filepath"

	"github.com/nivesh-jain/thobSyncDev.git/internal/minio"
	"github.com/spf13/cobra"
)

var uploadFileCmd = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload a file to a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucketName, _ := cmd.Flags().GetString("bucket")
		filePath, _ := cmd.Flags().GetString("file")

		if bucketName == "" || filePath == "" {
			log.Fatalln("Bucket name and file path are required.")
		}

		client := minio.NewClient()
		objectName := filepath.Base(filePath)

		minio.UploadFile(client, bucketName, objectName, filePath)
	},
}

func init() {
	uploadFileCmd.Flags().StringP("bucket", "b", "", "Name of the bucket")
	uploadFileCmd.Flags().StringP("file", "f", "", "Path to the file to upload")
	rootCmd.AddCommand(uploadFileCmd)
}
