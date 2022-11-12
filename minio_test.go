package main

import (
	"context"
	"log"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var bucketName = "go-mogu"

func getMinioClient() *minio.Client {
	endpoint := "minio-api.ithhit.cn"
	accessKeyID := "***"
	secretAccessKey := "***"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}

func TestMinio(t *testing.T) {
	ctx := context.Background()

	minioClient := getMinioClient()
	// Make a new bucket called mymusic.

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	//Upload the zip file
	objectName := "/test/NoLsp.zip"
	filePath := "./NoLsp.zip"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	minioClient := getMinioClient()

	// Make a new bucket called mymusic.
	// Upload the zip file with FPutObject
	err := minioClient.RemoveObject(ctx, bucketName, "1668268327205.jpg", minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

}
