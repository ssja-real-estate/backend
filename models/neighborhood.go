package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Neighborhood struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	MapInfo mapInfo            `json:"mapInfo" bsob:"mapInfo"`
}
