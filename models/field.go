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
	FiledValue interface{}        `json:"value" bson:"value"`
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
	}

}

func (filed *Field) Validate() error {
	if filed.Optional == true {
		return nil
	}
	if filed.FiledValue == nil {
		return errors.New(fmt.Sprint("فیلد ", filed.Title, " نباید خالی باشد"))
	}
	if filed.Fields != nil {
		for _, item := range filed.Fields {
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
	case Text:
		field.FiledValue = text{}
	case Number:
		field.FiledValue = number{}
	// case Select:
	// 	field.Value =
	case Bool:
		field.FiledValue = false
	case Conditional:
		{
			field.FiledValue = false
		}
	case Image:
		field.FiledValue = make([]string, 0)
	case Range:
		{
			field.FiledValue = arrayint{}
			field.Max = 0
			field.Min = 0
		}
	}
}
