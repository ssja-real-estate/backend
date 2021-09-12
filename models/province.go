package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Province struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Cities    []City        `json:"cites" bson:"cities"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
