package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type AssignmentType struct {
	// required: false
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	// required: false
	CreatedAt time.Time `json:"-" bson:"created_at" `
	// required: false
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
}
