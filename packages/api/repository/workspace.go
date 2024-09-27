package repository

import "go.mongodb.org/mongo-driver/mongo"

type WorkSpaceRepository interface {
}

type workspaceRepository struct {
	col *mongo.Collection
}

func NewWorkspaceRepository(
	col *mongo.Collection,
) UserRepository {
	return &userRepository{}
}
