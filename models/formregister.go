package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type FormRegister struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Form         Form          `json:"form" bson:"form"`
	RegisterDate string        `json:"registerDate" bson:"registerDate"`
	UserId       bson.ObjectId `json:"userId" bson:"userId"`
	Accept       bool          `json:"accept" bson:"accept"`
	CreatedAt    time.Time     `json:"-" bson:"createdAt"`
	UpdatedAt    time.Time     `json:"-" bson:"updatedAt"`
}
