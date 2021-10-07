package models

import "gopkg.in/mgo.v2/bson"

type Valuetype interface{}
type Field struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	// Name      string        `json:"name" bson:"name"`
	Title     string    `json:"title" bson:"title"`
	Value     Valuetype `json:"value" bson:"value"`
	Min       float64   `json:"min" bson:"min"`
	Max       float64   `json:"max" bson:"max"`
	Optional  bool      `json:"optional" bson:"optional"`
	Options   []string  `json:"option" bson:"option"`
	Fields    []Field   `json:"fileds" bson:"fileds"`
	Fieldtype string    `json:"fieldtype" bson:"fieldtype"`
}

func (field *Field) updateid() {
	field.Id = bson.NewObjectId()
	if len(field.Fields) > 0 {
		for i := 0; i < len(field.Fields); i++ {
			field.Fields[i].updateid()

		}
	}
}
