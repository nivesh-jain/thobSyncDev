package cmd

import (
	"context"
	"fmt"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var deleteBucketCmd = &cobra.Command{
	Use:   "delete-bucket",
	Short: "Delete a bucket (Admin Only)",
	Run: func(cmd *cobra.Command, args []string) {
		username := auth.CheckUserRole("Admin")

		client := db.GetMinioClient()
		err := client.RemoveBucket(context.Background(), bucketName)
		if err != nil {
			fmt.Printf("Failed to delete bucket: %v\n", err)
			db.LogOperation(username, "DeleteBucket", "", bucketName, "Failure")
			return
		}

		fmt.Println("Bucket deleted successfully!")
		db.LogOperation(username, "DeleteBucket", "", bucketName, "Success")
	},
}

func init() {
	deleteBucketCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket to delete")
	rootCmd.AddCommand(deleteBucketCmd)
}
