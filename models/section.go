package models

import "gopkg.in/mgo.v2/bson"

type Section struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	// Name   string        `json:"name" bson:"name"`
	Title  string  `json:"title" bson:"title"`
	Fileds []Field `json:"fields" bson:"fields"`
}

func (section *Section) updateid() {

	section.Id = bson.NewObjectId()
	for i := 0; i < len(section.Fileds); i++ {
		section.Fileds[i].updateid()

	}
}
