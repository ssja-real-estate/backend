package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignmentType struct {
	// required: false
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
