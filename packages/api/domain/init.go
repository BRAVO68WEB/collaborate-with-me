package domain

import (
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/repository"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Domain struct {
	db         db.Connection
	awsSession *session.Session
	repo       repository.Repository

	User      UserDomain
	Workspace WorkspaceDomain
}

func Init(
	db db.Connection,
	awsSession *session.Session,
) Domain {
	s := Domain{
		repo: repository.Init(
			db,
			awsSession,
		),
	}

	return s
}
