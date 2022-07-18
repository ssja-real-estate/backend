package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"password" bson:"password"`
	Mobile   string             `json:"mobile" bson:"mobile"`
	// owner 1 Admin 2 User 3
	Role       int       `json:"role"  bson:"role"`
	VerifyCode string    `json:"-" bson:"verifyCode"`
	Verify     bool      `json:"-" bson:"verify"`
	Credit     *Credit   `json:"credit" bson:"credit"`
	CreatedAt  time.Time `json:"-" bson:"createdAt"`
	UpdatedAt  time.Time `json:"-" bson:"updatedAt"`
}
