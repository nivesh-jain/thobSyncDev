package db

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

// GetMinioClient initializes and returns a MinIO client instance.
func GetMinioClient() *minio.Client {
	endpoint := viper.GetString("minio.endpoint")
	accessKey := viper.GetString("minio.accessKeyID")
	secretKey := viper.GetString("minio.secretAccessKey")
	useSSL := viper.GetBool("minio.useSSL")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	return client
}
