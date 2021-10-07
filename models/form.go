package models

import "gopkg.in/mgo.v2/bson"

type Form struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	// Name string        `json:"name" bson:"name"`
	// Sections []map[string]interface{} `json:"sections" bson:"sections"`
	Title    string    `json:"title" bson:"title"`
	Sections []Section `json:"sections" bson:"sections"`
}

func (form *Form) Updateid() {

	for i := 0; i < len(form.Sections); i++ {
		form.Sections[i].updateid()

	}
}
