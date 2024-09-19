package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/helpers"
	"github.com/aws/aws-sdk-go/aws/session"
)

type UploadResponse struct {
	is_success bool
	url        string
}

type UploadResponseMultiple struct {
	is_success bool
	urls       []string
}

type UploadDomain interface {
	SingleFileUpload(
		file graphql.Upload,
	) (string, error)
	MultipleFileUpload(
		files []graphql.Upload,
	) ([]string, error)
}

type uploadDomain struct {
	awsSession *session.Session
}

func NewUploadDomain(
	awsSession *session.Session,
) UploadDomain {
	return &uploadDomain{
		awsSession: awsSession,
	}
}

func (d *uploadDomain) SingleFileUpload(
	file graphql.Upload,
) (string, error) {

	bucket_name := os.Getenv("S3_BUCKET")
	current_ts := fmt.Sprintf("%d", time.Now().Unix())

	if !helpers.IsImage(file.Filename) {
		return "", fmt.Errorf("File is not an image")
	}

	upload, err := helpers.UploadFile(
		d.awsSession,
		&file.File,
		current_ts+"_"+file.Filename,
		bucket_name,
	)

	if err != nil {
		return "", err
	}

	return upload, nil
}

func (d *uploadDomain) MultipleFileUpload(
	files []graphql.Upload,
) ([]string, error) {
	var urls []string
	bucket_name := os.Getenv("S3_BUCKET")
	current_ts := fmt.Sprintf("%d", time.Now().Unix())

	for _, file := range files {
		upload, err := helpers.UploadFile(
			d.awsSession,
			&file.File,
			current_ts+"_"+file.Filename,
			bucket_name,
		)

		if err != nil {
			return nil, err
		}

		urls = append(urls, upload)
	}

	return urls, nil
}
