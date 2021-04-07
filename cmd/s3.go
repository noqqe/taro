package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(filename string) bool {
	key := aws.String(filename)
	bucket := aws.String(os.Getenv("B2_BUCKET"))
	appKeyId := os.Getenv("B2_APPLICATION_KEY_ID")
	appKey := os.Getenv("B2_APPLICATION_KEY")

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(appKeyId, appKey, ""),
		Endpoint:         aws.String("https://s3.us-west-001.backblazeb2.com"),
		Region:           aws.String("us-west-001"),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	// Upload a new object "testfile.txt" with the string "S3 Compatible API"
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("S3 Compatible API"),
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		fmt.Printf("Failed to upload object %s/%s, %s\n", *bucket, *key, err.Error())
		return false
	}

	// o, _ := s3Client.GetObject(&s3.GetObjectInput{
	// 	Bucket: bucket,
	// 	Key:    key,
	// })

	fmt.Printf("Successfully uploaded key %s\n", *key)
	return true
}
