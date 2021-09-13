package models

import "gopkg.in/mgo.v2/bson"

type Form struct {
	Id       bson.ObjectId            `json:"id" bson:"_id"`
	Name     string                   `json:"name" bson:"name"`
	Sections []map[string]interface{} `json:"sections" bson:"sections"`
}
