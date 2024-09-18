package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Role      string             `json:"role"`
	IsActive  bool               `json:"is_active"`
	CreatedAt primitive.DateTime `json:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at"`
}
