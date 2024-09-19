package domain

import "go.mongodb.org/mongo-driver/mongo"

type WorkspaceDomain interface {
}

type workspaceDomain struct {
}

func NewWorkspaceRepository(
	col *mongo.Collection,
) WorkspaceDomain {
	return &workspaceDomain{}
}
