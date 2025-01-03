package minio

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
)

// ListBuckets retrieves all buckets from the MinIO server
func ListBuckets(client *minio.Client) []minio.BucketInfo {
	buckets, err := client.ListBuckets(context.Background())
	if err != nil {
		log.Fatalf("Failed to list buckets: %v", err)
	}
	return buckets
}

// UploadFile uploads a file to the specified bucket
func UploadFile(client *minio.Client, bucketName, objectName, filePath string) {
	_, err := client.FPutObject(
		context.Background(),
		bucketName,
		objectName,
		filePath,
		minio.PutObjectOptions{},
	)
	if err != nil {
		log.Fatalf("Failed to upload file '%s' to bucket '%s': %v", filePath, bucketName, err)
	}
	fmt.Printf("Successfully uploaded '%s' to bucket '%s'\n", objectName, bucketName)
}

// DownloadFile downloads a file from the specified bucket
func DownloadFile(client *minio.Client, bucketName, objectName, destPath string) {
	err := client.FGetObject(
		context.Background(),
		bucketName,
		objectName,
		destPath,
		minio.GetObjectOptions{},
	)
	if err != nil {
		log.Fatalf("Failed to download file '%s' from bucket '%s': %v", objectName, bucketName, err)
	}
	fmt.Printf("Successfully downloaded '%s' from bucket '%s' to '%s'\n", objectName, bucketName, destPath)
}

// DeleteFile deletes a file from the specified bucket
func DeleteFile(client *minio.Client, bucketName, objectName string) {
	err := client.RemoveObject(
		context.Background(),
		bucketName,
		objectName,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		log.Fatalf("Failed to delete file '%s' from bucket '%s': %v", objectName, bucketName, err)
	}
	fmt.Printf("Successfully deleted '%s' from bucket '%s'\n", objectName, bucketName)
}

// ListFiles lists all files in the specified bucket
func ListFiles(client *minio.Client, bucketName string) {
	objectCh := client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})

	fmt.Printf("Files in bucket '%s':\n", bucketName)
	found := false
	for object := range objectCh {
		if object.Err != nil {
			log.Fatalf("Failed to list objects in bucket '%s': %v", bucketName, object.Err)
		}
		found = true
		fmt.Printf("- %s (Size: %d bytes)\n", object.Key, object.Size)
	}

	if !found {
		fmt.Printf("No files found in bucket '%s'\n", bucketName)
	}
}

// CreateBucket creates a new bucket in the MinIO server
func CreateBucket(client *minio.Client, bucketName string) {
	err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		log.Fatalf("Failed to create bucket '%s': %v", bucketName, err)
	}
	fmt.Printf("Successfully created bucket '%s'\n", bucketName)
}

// DeleteBucket deletes a bucket from the MinIO server
func DeleteBucket(client *minio.Client, bucketName string) {
	err := client.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		log.Fatalf("Failed to delete bucket '%s': %v", bucketName, err)
	}
	fmt.Printf("Successfully deleted bucket '%s'\n", bucketName)
}
