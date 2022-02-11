package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type City struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	MapInfo       mapInfo            `json:"mapinfo" bson:"mapinfo"`
	Neighborhoods []Neighborhood     `json:"neighborhoods" bson:"neighborhoods"`
}
