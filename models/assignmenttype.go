package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignmentType struct {
	// required: false
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	// required: false
	CreatedAt time.Time `json:"-" bson:"createdAt" `
	// required: false
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}
