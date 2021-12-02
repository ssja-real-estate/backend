package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// swagger:model
type EstateType struct {
	// required: false
	Id bson.ObjectId `json:"id" bson:"_id"`
	// The name for a EstateType
	// example: خرید
	// required: true
	Name string `json:"name" bson:"name"`
	// required: false
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	// required: false
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}
