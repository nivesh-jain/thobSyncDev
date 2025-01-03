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

var uploadFileCmd = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload a file to a specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucketName, _ := cmd.Flags().GetString("bucket")
		filePath, _ := cmd.Flags().GetString("file")

		if bucketName == "" || filePath == "" {
			log.Fatalln("Bucket name and file path are required.")
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

		fileName := filepath.Base(filePath)
		contentType := "application/octet-stream"

		_, err = minioClient.FPutObject(context.Background(), bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Successfully uploaded %s to bucket %s\n", fileName, bucketName)
	},
}

func init() {
	uploadFileCmd.Flags().StringP("bucket", "b", "", "Name of the bucket")
	uploadFileCmd.Flags().StringP("file", "f", "", "Path to the file to upload")
}
