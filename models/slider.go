package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Slider struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Path string             `json:"path"`
}
