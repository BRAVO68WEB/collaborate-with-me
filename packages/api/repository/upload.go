package repository

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go/aws/session"
)

type UploadRepository interface {
	UploadFile(
		file graphql.Upload,
	) (string, error)
}

type uploadRepository struct {
	awsSession *session.Session
}

func NewUploadRepository(
	awsSession *session.Session,
) UploadRepository {
	return &uploadRepository{
		awsSession: awsSession,
	}
}

func (r *uploadRepository) UploadFile(
	file graphql.Upload,
) (string, error) {
	return "", nil
}
