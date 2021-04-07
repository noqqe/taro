package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func buildS3Client() *s3.S3 {
	appKeyId := os.Getenv("B2_APPLICATION_KEY_ID")
	appKey := os.Getenv("B2_APPLICATION_KEY")

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(appKeyId, appKey, ""),
		Endpoint:         aws.String("https://s3.us-west-001.backblazeb2.com"),
		Region:           aws.String("us-west-001"),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)
	return s3.New(newSession)
}

func UploadToS3(filename string, in string) bool {

	bucket := aws.String(os.Getenv("B2_BUCKET"))

	fmt.Printf("Uploading %s to S3 Bucket %s\n", in, *bucket)

	file, err := os.Open(in)
	if err != nil {
		fmt.Printf("Failed to open file %s\n%s\n", in, err.Error())
		return false
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Upload a new object "testfile.txt" with the string "S3 Compatible API"
	s3Client := buildS3Client()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:             bucket,
		Key:                aws.String(filename),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(size),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})
	if err != nil {
		fmt.Printf("Failed to upload object %s/%s, %s\n", *bucket, filename, err.Error())
		return false
	}

	fmt.Printf("Successfully uploaded key %s\n", filename)
	return true
}

func GetPhotoFromS3(name string) io.Reader {

	bucket := aws.String(os.Getenv("B2_BUCKET"))
	s3Client := buildS3Client()
	o, _ := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    aws.String(name),
	})
	return o.Body
}
