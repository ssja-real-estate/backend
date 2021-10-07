package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Password string        `json:"password" bson:"password"`
	Mobile   string        `json:"mobile" bson:"mobile"`
	// owner 0 Admin 1 User 2
	Role      int       `json:"role" bson:"role"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
}
