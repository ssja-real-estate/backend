package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Unit struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"-" bson:"createdAt"`
	UpdatedAt time.Time          `json:"-" bson:"updatedAt"`
}
