package models

import "gopkg.in/mgo.v2/bson"

type Form struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Title            string        `json:"title" bson:"title"`
	AssignmentTypeID bson.ObjectId `json:"assignment_type_id" bson:"assignment_type_id"`
	EstateTypeID     bson.ObjectId `json:"estate_type_id" bson:"estate_type_id"`
	Sections         []Section     `json:"sections" bson:"sections"`
}

func (form *Form) Updateid() {

	for i := 0; i < len(form.Sections); i++ {
		form.Sections[i].updateid()

	}
}
