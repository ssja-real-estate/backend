package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HeadFilter struct {
	AssignmentTypeID primitive.ObjectID `json:"assignmentTypeId" bson:"assignmentTypeId"`
	EstateTypeID     primitive.ObjectID `json:"estateTypeId" bson:"estateTypeId"`
	ProvinceID       primitive.ObjectID `json:"provicneId" bson:"provinceId"`
	CityID           primitive.ObjectID `json:"cityId" bson:"cityId"`
	NeighborhoodID   primitive.ObjectID `json:"neighbordhoodId" bson:"neighbordhoodId"`
}
