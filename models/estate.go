package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Estate struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	UserId         primitive.ObjectID `json:"userId" bson:"userId"`
	Verified       bool               `json:"verified" bson:"verified"`
	Latitude       float64            `json:"latitude" bson:"latitude"`
	Longitude      float64            `json:"longitude" bson:"longitude"`
	CityId         primitive.ObjectID `json:"cityId" bson:"cityId"`
	ProvinceId     primitive.ObjectID `json:"provinceId" bson:"provinceId"`
	NeighborhoodId primitive.ObjectID `json:"neighborhoodId" bson:"neighborhoodId"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdateAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
	DataForm       Form               `json:"dataForm" bson:"dataForm"`
}
