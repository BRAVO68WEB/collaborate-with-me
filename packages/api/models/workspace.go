package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workspace struct {
	ID                primitive.ObjectID   `json:"id"  bson:"_id,omitempty"`
	Name              string               `json:"name"`
	IsActive          bool                 `json:"is_active"`
	IsPublic          bool                 `json:"is_public"`
	Owner             primitive.ObjectID   `json:"owner"`
	Collaborators     []primitive.ObjectID `json:"collaborators"`
	ExcalidrawObjects []ExcaliObjects      `json:"excalidraw_objects" bson:"excalidraw_objects"`
	CreatedAt         primitive.DateTime   `json:"created_at"`
	UpdatedAt         primitive.DateTime   `json:"updated_at"`
}
