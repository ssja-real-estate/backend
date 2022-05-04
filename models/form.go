package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Form struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Title            string             `json:"title" bson:"title"`
	AssignmentTypeID primitive.ObjectID `json:"assignmentTypeId" bson:"assignmentTypeId"`
	EstateTypeID     primitive.ObjectID `json:"estateTypeId" bson:"estateTypeId"`
	//  Sections         []Section          `json:"sections" bson:"sections"`
	Fields []Field `json:"fields" bson:"fields"`
}

func (form *Form) Updateid() {
	for i := 0; i < len(form.Fields); i++ {
		form.Fields[i].updateid()
	}
}
func (form *Form) Validate() error {
	for _, items := range form.Fields {
		err := items.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
