package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Province struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	MapInfo   mapInfo            `json:"mapInfo" bson:"mapInfo"`
	Cities    []City             `json:"cities" bson:"cities"`
	CreatedAt time.Time          `json:"-" bson:"createdAt"`
	UpdatedAt time.Time          `json:"-" bson:"updatedAt"`
}
