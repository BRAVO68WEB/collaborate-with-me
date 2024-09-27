package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExcaliObjects struct {
	ID               primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	ExcalidrawObject interface{}        `json:"excalidraw_object" bson:"excalidraw_object"`
	CreatedAt        primitive.DateTime `json:"created_at"`
	UpdatedAt        primitive.DateTime `json:"updated_at"`
}
