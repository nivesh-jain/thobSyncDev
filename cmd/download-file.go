package cmd

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
)

var downloadFileCmd = &cobra.Command{
	Use:   "download-file",
	Short: "Download a file from a specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucketName, _ := cmd.Flags().GetString("bucket")
		objectName, _ := cmd.Flags().GetString("object")
		destination, _ := cmd.Flags().GetString("dest")

		if bucketName == "" || objectName == "" || destination == "" {
			log.Fatalln("Bucket name, object name, and destination path are required.")
		}

		endpoint := "localhost:9000"
		accessKeyID := "Gx0S3h31P8SfmOWhm3Tg"
		secretAccessKey := "XAqfnX6Q77PhtEUhyjziZj8bsPpz9PoSLtgSh1yY"
		useSSL := false

		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			log.Fatalln(err)
		}

		destPath := filepath.Join(destination, objectName)

		err = minioClient.FGetObject(context.Background(), bucketName, objectName, destPath, minio.GetObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Successfully downloaded %s to %s\n", objectName, destPath)
	},
}

func init() {
	downloadFileCmd.Flags().StringP("bucket", "b", "", "Name of the bucket")
	downloadFileCmd.Flags().StringP("object", "o", "", "Name of the object to download")
	downloadFileCmd.Flags().StringP("dest", "d", "", "Destination directory path")
}