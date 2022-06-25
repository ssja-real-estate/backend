package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Section struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Title  string             `json:"title" bson:"title"`
	Fileds []Field            `json:"fields" bson:"fields"`
}

func (section *Section) updateid() {
	section.Id = primitive.NewObjectID()
	for i := 0; i < len(section.Fileds); i++ {
		section.Fileds[i].updateid()
	}
}

func (section *Section) Validate() error {

	for _, item := range section.Fileds {
		err := item.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
