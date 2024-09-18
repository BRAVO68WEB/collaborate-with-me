package repository

import (
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Repository struct {
	db         db.Connection
	awsSession session.Session

	User      UserRepository
	Workspace WorkSpaceRepository
}

func Init(
	db db.Connection,
	awsSession *session.Session,
) *Repository {
	s := Repository{
		User: NewUserRepository(
			db.GetCollection("users"),
		),
		Workspace: NewWorkspaceRepository(
			db.GetCollection("workspaces"),
		),
	}

	return &s
}
