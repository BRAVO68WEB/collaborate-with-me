package helpers

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/h2non/bimg"
)

func OptimizeImage(file *io.ReadSeeker) (*bytes.Reader, error) {
	buffer, err := io.ReadAll(*file)
	if err != nil {
		panic(err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	return bytes.NewReader(newImage), err
}

func UploadFile(awsSession *session.Session, file *io.ReadSeeker, name string, bucketName ...string) (string, error) {
	uploader := s3manager.NewUploader(awsSession)

	image, err := OptimizeImage(file)
	if err != nil {
		return "", err
	}
	println("hello 2")
	bucket := os.Getenv("S3_BUCKET")
	if len(bucketName) > 0 {
		bucket = bucketName[0]
	}

	key := fmt.Sprintf("images/%s%s", name, ".webp")
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(key),
		Body:   image,
	})
	if err != nil {
		return "", err
	}
	imgUrl := fmt.Sprintf("%s/%s", os.Getenv("S3_OBJECT_URL"), key)
	return imgUrl, nil
}
