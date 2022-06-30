package repository

import (
	"context"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const paymentCollection = "payments"

type PaymentRepository interface {
	Save(payment *models.Payment) error
	Update(payment *models.Payment, paymentid primitive.ObjectID) error
	GetPayments() ([]models.Payment, error)
	DeletePayment(paymentid primitive.ObjectID) error
	GetPaymentByID(paymentid primitive.ObjectID) (models.Payment, error)
	GetPaymentByCreditDuration(credit int, duration int) (bool, error)
}

type paymentRepository struct {
	c *mongo.Collection
}

func NewPaymentRepository(DB *mongo.Client) PaymentRepository {
	return &paymentRepository{db.GetCollection(db.DB, paymentCollection)}
}

func (r *paymentRepository) Save(payment *models.Payment) error {
	_, err := r.c.InsertOne(context.TODO(), payment)
	return err
}

func (r *paymentRepository) Update(payment *models.Payment, paymentid primitive.ObjectID) error {
	filter := bson.M{"_id": paymentid}
	update := bson.M{"$set": &payment}
	_, err := r.c.UpdateOne(context.TODO(), filter, update)
	return err
}
func (r *paymentRepository) GetPayments() ([]models.Payment, error) {

	var paymens []models.Payment
	result, err := r.c.Find(context.TODO(), bson.D{})
	if err != nil {
		return []models.Payment{}, err

	}
	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var payment models.Payment
		if err = result.Decode(&payment); err != nil {
			return []models.Payment{}, err
		}
		paymens = append(paymens, payment)
	}
	if paymens != nil {
		return paymens, nil
	}
	return []models.Payment{}, nil
}

func (r *paymentRepository) DeletePayment(paymnetid primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": paymnetid})
	return err

}
func (r *paymentRepository) GetPaymentByID(paymentid primitive.ObjectID) (models.Payment, error) {
	var payment models.Payment
	err := r.c.FindOne(context.TODO(), bson.M{"_id": paymentid}).Decode(&payment)
	return payment, err
}
func (r *paymentRepository) GetPaymentByCreditDuration(creidt int, duration int) (bool, error) {
	var payment models.Payment
	err := r.c.FindOne(context.TODO(), bson.M{"credit": creidt, "duration": duration}).Decode(&payment)
	if err != nil {
		return false, err
	}
	if payment.Id.IsZero() {
		return false, nil
	}
	return true, nil
}
