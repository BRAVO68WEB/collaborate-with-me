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
		db:         db,
		awsSession: awsSession,
		repo: repository.Init(
			db,
			awsSession,
		),
		User:      NewUserDomain(db),
		Workspace: NewWorkspaceRepository(db.GetCollection("workspace")),
	}

	return s
}
