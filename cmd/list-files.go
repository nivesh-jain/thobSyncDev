package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listFilesCmd = &cobra.Command{
	Use:   "list-files",
	Short: "List all files in a specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucketName, _ := cmd.Flags().GetString("bucket")

		if bucketName == "" {
			log.Fatalln("Bucket name is required.")
		}

		endpoint := viper.GetString("minio.endpoint")
		accessKeyID := viper.GetString("minio.accessKeyID")
		secretAccessKey := viper.GetString("minio.secretAccessKey")
		useSSL := viper.GetBool("minio.useSSL")

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
