package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	Caption   string             `json:"caption,omitempty" bson:"caption,omitempty"`
	Timestamp string             `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
