package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Unit struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	CreatedAt time.Time     `json:"- bson:"created_at"`
	UpdatedAt time.Time     `json:"-" bson:"updated_at"`
}
