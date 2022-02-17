package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Valuetype interface {
}
type text struct {
	value string
}
type number struct {
	value int
}
type boolean struct {
	value bool
}
type arraystring struct {
	value []string
}
type arrayint struct {
	value []int
}

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
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	FieldValue interface{}        `json:"value" bson:"value"`
	Min        float64            `json:"min" bson:"min"`
	Max        float64            `json:"max" bson:"max"`
	Optional   bool               `json:"optional" bson:"optional"`
	Options    []string           `json:"options" bson:"options"`
	Fields     []Field            `json:"fields" bson:"fields"`
	Type       int                `json:"type"  bson:"type"`
}

// to do set value from type by enum
func (field *Field) updateid() {
	field.Id = primitive.NewObjectID()
	if len(field.Fields) > 0 {
		for i := 0; i < len(field.Fields); i++ {
			field.Fields[i].updateid()
			field.setValue()
		}
	} else {
		field.setValue()
	}
}

func (field *Field) Validate() error {
	if field.Optional == true {
		return nil
	}
	if field.FieldValue == nil {
		return errors.New(fmt.Sprint("فیلد ", field.Title, " نباید خالی باشد"))
	}
	if field.Fields != nil {
		for _, item := range field.Fields {
			err := item.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (field *Field) setValue() {
	switch field.Type {
	// case Text:
	// 	field.FieldValue = ""
	// case Number:
	// 	field.FieldValue = 0
	//   case Select:
	// 	field.FieldValue =boolean{}
	case Bool:
		field.FieldValue = false
	case Conditional:
		{
			field.FieldValue = false
		}
	case Image:
		field.FieldValue = make([]string, 0)
	case Range:
		{
			field.FieldValue = arrayint{}
			field.Max = 0
			field.Min = 0
		}
	}
}
