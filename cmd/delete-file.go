package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nivesh-jain/thobSyncDev.git/internal/rbac"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteFileCmd = &cobra.Command{
	Use:   "delete-file",
	Short: "Delete a file from a specified bucket",
	Run: func(cmd *cobra.Command, args []string) {
		role, _ := cmd.Flags().GetString("role")
		bucketName, _ := cmd.Flags().GetString("bucket")
		objectName, _ := cmd.Flags().GetString("object")

		if role == "" {
			log.Fatalln("Role is required. Use the --role flag to specify a role (e.g., Admin, Editor, Viewer).")
		}

		if bucketName == "" || objectName == "" {
			log.Fatalln("Bucket name and object name are required.")
		}

		endpoint := viper.GetString("minio.endpoint")
		accessKeyID := viper.GetString("minio.accessKeyID")
		secretAccessKey := viper.GetString("minio.secretAccessKey")
		useSSL := viper.GetBool("minio.useSSL")

		// Check permission
		if err := rbac.CheckPermission(role, "delete-file"); err != nil {
			log.Fatalf("Access Denied: %v\n", err)
		}

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
	deleteFileCmd.Flags().StringP("role", "r", "", "Role of the user (Admin, Editor, Viewer)")
	deleteFileCmd.Flags().StringP("bucket", "b", "", "Name of the bucket")
	deleteFileCmd.Flags().StringP("object", "o", "", "Name of the object to delete")

}
