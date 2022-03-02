package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Estate struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	Verified     bool               `json:"verified" bson:"verified"`
	Latitude     float64            `json:"latitude" bson:"latitude"`
	Longitude    float64            `json:"longitude" bson:"longitude"`
	City         EstateLocation     `json:"city" bson:"city"`
	Province     EstateLocation     `json:"province" bson:"province"`
	Neighborhood EstateLocation     `json:"neighborhood" bson:"neighborhood"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdateAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	DataForm     Form               `json:"dataForm" bson:"dataForm"`
}
