package models

import (
	"errors"
	"fmt"
	"realstate/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title" bson:"title"`
	Credit   int                `json:"credit" bson:"credit"`
	Duration int                `json:"duration" bson:"duration"`
}

func (payment *Payment) Validate() error {
	if len(payment.Title) < 3 {
		return errors.New(fmt.Sprint("فیلد", payment.Title, "نباید خالی باشد"))
	}
	if payment.Credit < 0 {
		return util.ErrCredete
	}
	if payment.Duration < 0 {
		return util.ErrDuration
	}
	return nil
}
