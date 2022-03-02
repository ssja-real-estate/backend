package repository

import (
	"context"
	"os"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const estateCollections = "estate"

type EstateRepository interface {
	SaveEstate(estate *models.Estate) error
	UpdateEstate(estate *models.Estate, estateid primitive.ObjectID) error
	DeleteEstate(estateid primitive.ObjectID) error
	GetEstateById(estateid primitive.ObjectID) (models.Estate, error)
	GetEstateNotVerified() ([]models.Estate, error)
	GetEstateVerified() ([]models.Estate, error)
	VerifyEstated(estaeid primitive.ObjectID) (int, error)
	GetEstateByUserID(userId primitive.ObjectID) ([]models.Estate, error)
}

type estateRepository struct {
	c *mongo.Collection
}

func NewEstateRepository(DB *mongo.Client) EstateRepository {
	return &estateRepository{db.GetCollection(DB, estateCollections)}
}

func (r *estateRepository) SaveEstate(estate *models.Estate) error {
	_, err := r.c.InsertOne(context.TODO(), estate)
	return err
}

func (r *estateRepository) UpdateEstate(estate *models.Estate, estateid primitive.ObjectID) error {

	filter := bson.M{"_id": estateid}
	update := bson.M{"$set": &estate}
	_, err := r.c.UpdateOne(context.TODO(), filter, update)
	return err

}
func (r *estateRepository) DeleteEstate(estateid primitive.ObjectID) error {
	err := os.RemoveAll(estateid.Hex())
	if err != nil {
		return err
	}
	_, err = r.c.DeleteOne(context.TODO(), bson.M{"_id": estateid})

	return err
}

func (r *estateRepository) GetEstateById(estateid primitive.ObjectID) (models.Estate, error) {
	var estate models.Estate
	err := r.c.FindOne(context.TODO(), bson.M{"_id": estateid}).Decode(&estate)
	if err != nil {
		return models.Estate{}, err
	}
	return estate, err
}
func (r *estateRepository) GetEstateVerified() ([]models.Estate, error) {

	var estates []models.Estate
	result, err := r.c.Find(context.TODO(), bson.M{"verified": true})
	if err != nil {
		return []models.Estate{}, err
	}

	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var estate models.Estate
		if err = result.Decode(&estate); err != nil {
			return []models.Estate{}, err
		}
		estates = append(estates, estate)

	}
	return estates, nil

}
func (r *estateRepository) VerifyEstated(estaeid primitive.ObjectID) (int, error) {
	query := bson.M{"_id": estaeid}
	update := bson.M{"$set": bson.M{"verified": true}}
	res, err := r.c.UpdateOne(context.TODO(), query, update)
	return int(res.ModifiedCount), err

}
func (r *estateRepository) GetEstateNotVerified() ([]models.Estate, error) {
	var estates []models.Estate
	result, err := r.c.Find(context.TODO(), bson.M{"verified": false})
	if err != nil {
		return []models.Estate{}, err
	}

	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var estate models.Estate
		if err = result.Decode(&estate); err != nil {
			return []models.Estate{}, err
		}
		estates = append(estates, estate)

	}
	return estates, nil

}

func (r *estateRepository) GetEstateByUserID(userId primitive.ObjectID) ([]models.Estate, error) {
	query := bson.M{"userId": userId}
	var estates []models.Estate
	result, err := r.c.Find(context.TODO(), query)
	if err != nil {
		return []models.Estate{}, err
	}
	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var estate models.Estate
		if err = result.Decode(&estate); err != nil {
			return []models.Estate{}, err
		}
		estates = append(estates, estate)
	}
	return estates, nil
}
