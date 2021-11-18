package models

import "gopkg.in/mgo.v2/bson"

type Valuetype interface{}

const (
	Text        = 0
	Number      = 1
	Select      = 2
	Bool        = 3
	Conditional = 4
	Image       = 5
	Range       = 6
)

type Field struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	// Name      string        `json:"name" bson:"name"`
	Title    string    `json:"title" bson:"title"`
	Value    Valuetype `json:"value" bson:"value"`
	Min      float64   `json:"min" bson:"min"`
	Max      float64   `json:"max" bson:"max"`
	Optional bool      `json:"optional" bson:"optional"`
	Options  []string  `json:"option" bson:"option"`
	Fields   []Field   `json:"fileds" bson:"fileds"`
	Type     int       `json:"type"  bson:"type"`
}

// to do set value from type by enum
func (field *Field) updateid() {
	field.Id = bson.NewObjectId()
	if len(field.Fields) > 0 {
		for i := 0; i < len(field.Fields); i++ {
			field.Fields[i].updateid()
			field.setValue()
		}
	}
}

func (field *Field) setValue() {
	switch field.Type {
	case Text:
		field.Value = ""
	case Number:
		field.Value = 0
	// case Select:
	// 	field.Value =
	case Bool:
		field.Value = false
	case Conditional:
		{
			field.Value = ""
		}

	case Image:
		field.Value = ""
	case Range:
		{
			field.Value = 0
			field.Max = 0
			field.Min = 0
		}
	}

}
