package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Estate struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	UserId     primitive.ObjectID `json:"userId" bson:"_userId"`
	Verified   bool               `json:"verified" bson:"verified"`
	Latitude   float64            `json:"latitude" bson:"latitude"`
	Langitude  float64            `json:"langitude" bson:"langitude"`
	CityId     primitive.ObjectID `json:"cityId" bson:"_cityId"`
	ProvinceId primitive.ObjectID `json:"provinceId" bson:"_provinceId"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdateAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
	DataForm   Form               `json:"dataForm" bson:"dataForm"`
}
