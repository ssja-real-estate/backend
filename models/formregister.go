package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstateRegister struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Form         Form               `json:"form" bson:"form"`
	RegisterDate string             `json:"registerDate" bson:"registerDate"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	Accept       bool               `json:"accept" bson:"accept"`
	CreatedAt    time.Time          `json:"-" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"-" bson:"updatedAt"`
}
