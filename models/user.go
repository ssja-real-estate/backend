package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	UserName  string        `json:"user_name" bson:"user_name"`
	Password  string        `json:"password" bson:"password"`
	Mobile    string        `json:"mobile" bson:"mobile"`
	Role      string        `json:"role" bson:"role"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
