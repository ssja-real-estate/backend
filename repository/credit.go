package repository

import (
	"context"
	"errors"
	"fmt"
	"realstate/db"
	"realstate/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const creditCollection = "credit"

type CreditRepository interface {
	Save(models.Credit) error
	Delete(primitive.ObjectID) error
	GetCredit(id primitive.ObjectID) (models.Credit, error)
}

type creditRepository struct {
	c *mongo.Collection
}

func NewCreditRepository(DB *mongo.Client) CreditRepository {
	return &creditRepository{db.GetCollection(db.DB, creditCollection)}
}

func (r *creditRepository) Save(credit models.Credit) error {
	_, err := r.c.InsertOne(context.TODO(), credit)
	return err
}

func (r *creditRepository) Delete(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
func (r *creditRepository) GetCredit(userid primitive.ObjectID) (models.Credit, error) {

	var credit models.Credit
	fmt.Println(userid)
	err := r.c.FindOne(context.TODO(), bson.M{"userId": userid}).Decode(&credit)
	if err != nil {
		fmt.Println(err)
		return models.Credit{}, err
	}
	fmt.Println(credit)
	t := time.Now()
	dif := t.Sub(credit.RegisterDate)
	duration := credit.RemainingDuration - int(dif.Hours())/24
	if duration > 0 {
		credit.Duration = duration
	} else {
		credit.Duration = 0
		return models.Credit{}, errors.New("credit is expire")
	}
	return credit, nil
}
