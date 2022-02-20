package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// swagger:model
type EstateType struct {
	// required: false
	Id primitive.ObjectID `json:"id" bson:"_id"`
	// The name for a EstateType
	// example: خرید
	// required: true
	Name string `json:"name" bson:"name"`
}
