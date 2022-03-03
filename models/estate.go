package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Estate struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	Estatetatus  EstateStatus       `json:"estateStatus" bson:"estateStatus"`
	Position     Position           `json:"position" bson:"position"`
	City         EstateLocation     `json:"city" bson:"city"`
	Province     EstateLocation     `json:"province" bson:"province"`
	Neighborhood EstateLocation     `json:"neighborhood" bson:"neighborhood"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdateAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	DataForm     Form               `json:"dataForm" bson:"dataForm"`
}
