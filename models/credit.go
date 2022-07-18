package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credit struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	UserId            primitive.ObjectID `json:"userId" bson:"userId"`
	RegisterDate      time.Time          `json:"registerDate" bson:"registerDate"`
	Duration          int                `json:"duration" bson:"duration"`
	RemainingDuration int                `json:"remainingDuration" bson:"remainingDuration"`
	Credit            int                `json:"credit" bson:"credit"`
	Title             string             `json:"title" bson:"title"`
}
