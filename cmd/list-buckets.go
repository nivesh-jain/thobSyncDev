// In cmd/list-buckets.go
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

var listBucketsCmd = &cobra.Command{
	Use:   "list-buckets",
	Short: "List all S3 buckets",
	Run: func(cmd *cobra.Command, args []string) {
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

		// List buckets
		buckets, err := minioClient.ListBuckets(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		for _, bucket := range buckets {
			fmt.Println(bucket.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listBucketsCmd)
}
