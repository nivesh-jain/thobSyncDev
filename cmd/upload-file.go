package cmd

// import (
// 	"log"
// 	"path/filepath"

// 	"github.com/nivesh-jain/thobSyncDev.git/internal/minio"
// 	"github.com/nivesh-jain/thobSyncDev.git/internal/rbac"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// var uploadFileCmd = &cobra.Command{
// 	Use:   "upload-file",
// 	Short: "Upload a file to a bucket",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		role := viper.GetString("role")
// 		if role == "" {
// 			log.Fatalln("You must initialize the CLI with 'init' before using this command.")
// 		}

// 		if err := rbac.CheckPermission(role, "upload-file"); err != nil {
// 			log.Fatalf("Access Denied: %v\n", err)
// 		}

// 		bucketName, _ := cmd.Flags().GetString("bucket")
// 		filePath, _ := cmd.Flags().GetString("file")

// 		if bucketName == "" || filePath == "" {
// 			log.Fatalln("Bucket name and file path are required.")
// 		}

// 		client := minio.NewClient()
// 		objectName := filepath.Base(filePath)

// 		minio.UploadFile(client, bucketName, objectName, filePath)
// 	},
// }

// func init() {
// 	uploadFileCmd.Flags().StringP("bucket", "b", "", "Name of the bucket")
// 	uploadFileCmd.Flags().StringP("file", "f", "", "Path to the file to upload")
// 	rootCmd.AddCommand(uploadFileCmd)
// }
