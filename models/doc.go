package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" `
	Path  string             `json:"path"`
	Type  int                `json:"type"`
}
