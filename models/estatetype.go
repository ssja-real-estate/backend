package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstateType struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Order int32              `json:"order" bson:"order"`
}
