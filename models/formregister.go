package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type FormRegister struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Form         Form          `json:"form" bson:"form"`
	RegisterDate string        `json:"register_date" bson:"register_date"`
	UserId       bson.ObjectId `json:"user_id" bson:"user_id"`
	Accept       bool          `json:"accept" bson:"accept"`
	CreatedAt    time.Time     `json:"-" bson:"created_at"`
	UpdatedAt    time.Time     `json:"-" bson:"updated_at"`
}
