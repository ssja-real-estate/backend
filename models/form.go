package models

import "gopkg.in/mgo.v2/bson"

type Form struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Title            string        `json:"title" bson:"title"`
	AssignmentTypeID bson.ObjectId `json:"assignmentTypeId" bson:"assignmentTypeId"`
	EstateTypeID     bson.ObjectId `json:"estateTypeId" bson:"estateTypeId"`
	Sections         []Section     `json:"sections" bson:"sections"`
}

func (form *Form) Updateid() {

	for i := 0; i < len(form.Sections); i++ {
		form.Sections[i].updateid()

	}
}
