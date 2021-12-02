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
	// owner 1 Admin 2 User 3
	Role       int       `json:"role"  bson:"role"`
	VerifyCode string    `json:"-" bson:"verifyCode"`
	Verify     bool      `json:"-" bson:"verify"`
	CreatedAt  time.Time `json:"-" bson:"createdAt"`
	UpdatedAt  time.Time `json:"-" bson:"updatedAt"`
}
