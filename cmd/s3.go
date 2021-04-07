package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(filename string, in string) bool {

	key := aws.String(filename)
	bucket := aws.String(os.Getenv("B2_BUCKET"))
	appKeyId := os.Getenv("B2_APPLICATION_KEY_ID")
	appKey := os.Getenv("B2_APPLICATION_KEY")

	file, _ := os.Open(in)
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

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
		Bucket:             bucket,
		Key:                key,
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(size),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
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
