package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
)

var deleteFileCmd = &cobra.Command{
	Use:   "delete-file",
	Short: "Delete a file from a specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucketName, _ := cmd.Flags().GetString("bucket")
		objectName, _ := cmd.Flags().GetString("object")

		if bucketName == "" || objectName == "" {
			log.Fatalln("Bucket name and object name are required.")
		}

		endpoint := "localhost:9000"
		accessKeyID := "Gx0S3h31P8SfmOWhm3Tg"
		secretAccessKey := "XAqfnX6Q77PhtEUhyjziZj8bsPpz9PoSLtgSh1yY"
		useSSL := false

		// Initialize MinIO client
		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			log.Fatalln(err)
		}

		// Delete the object
		err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Successfully deleted %s from bucket %s\n", objectName, bucketName)
	},
}

func init() {
	deleteFileCmd.Flags().StringP("bucket", "b", "", "Name of the bucket")
	deleteFileCmd.Flags().StringP("object", "o", "", "Name of the object to delete")
}
