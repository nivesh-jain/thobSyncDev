package cmd

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var createBucketCmd = &cobra.Command{
	Use:   "create-bucket",
	Short: "Create a new bucket (Admin Only)",
	Run: func(cmd *cobra.Command, args []string) {
		username := auth.CheckUserRole("Admin")

		client := db.GetMinioClient()
		err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Printf("Failed to create bucket: %v\n", err)
			db.LogOperation(username, "CreateBucket", "", bucketName, "Failure")
			return
		}

		fmt.Println("Bucket created successfully!")
		db.LogOperation(username, "CreateBucket", "", bucketName, "Success")
	},
}

func init() {
	createBucketCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket to create")
	rootCmd.AddCommand(createBucketCmd)
}
