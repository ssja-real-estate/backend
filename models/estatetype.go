package models

import (
	"time"

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
	// required: false
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	// required: false
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}
