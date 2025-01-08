package cmd

import (
	"context"
	"fmt"

	"github.com/nivesh-jain/thobSyncDev.git/internal/auth"
	"github.com/nivesh-jain/thobSyncDev.git/internal/db"
	"github.com/spf13/cobra"
)

var listBucketsCmd = &cobra.Command{
	Use:   "list-buckets",
	Short: "List all buckets (Admin Only)",
	Run: func(cmd *cobra.Command, args []string) {
		username := auth.CheckUserRole("Admin")

		client := db.GetMinioClient()
		buckets, err := client.ListBuckets(context.Background())
		if err != nil {
			fmt.Printf("Failed to list buckets: %v\n", err)
			db.LogOperation(username, "ListBuckets", "", "", "Failure")
			return
		}

		fmt.Println("Buckets:")
		for _, bucket := range buckets {
			fmt.Println(bucket.Name)
		}
		db.LogOperation(username, "ListBuckets", "", "", "Success")
	},
}

func init() {
	rootCmd.AddCommand(listBucketsCmd)
}
