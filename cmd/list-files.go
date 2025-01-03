package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
)

var listFilesCmd = &cobra.Command{
	Use:   "list-files",
	Short: "List all files in a specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucketName, _ := cmd.Flags().GetString("bucket")

		if bucketName == "" {
			log.Fatalln("Bucket name is required.")
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

		// List objects in the bucket
		objectCh := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})

		fmt.Printf("Files in bucket %s:\n", bucketName)
		for object := range objectCh {
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			fmt.Printf("- %s (Size: %d bytes)\n", object.Key, object.Size)
		}
	},
}

func init() {
	listFilesCmd.Flags().StringP("bucket", "b", "", "Name of the bucket to list files")
}
