package repository

import (
	"context"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const unitcollection = "units"

type UnitRepository interface {
	SaveUnit(unit *models.Unit) error
	UpdateUnit(unit *models.Unit) error
	GetUnitById(id primitive.ObjectID) (unit *models.Unit, err error)
	GetUnitByName(name string) (unit *models.Unit, err error)
	GetUnitAll() (units []models.Unit, err error)
	DeleteUnit(id primitive.ObjectID) error
}

type unitRepository struct {
	c *mongo.Collection
}

func NewUnitRepository(DB *mongo.Client) UnitRepository {
	return &unitRepository{db.GetCollection(db.DB, unitcollection)}
}

func (r *unitRepository) SaveUnit(unit *models.Unit) error {
	_, err := r.c.InsertOne(context.TODO(), unit)
	return err
}
func (r *unitRepository) UpdateUnit(unit *models.Unit) error {
	_, err := r.c.UpdateOne(context.TODO(), bson.M{"_id": unit.Id}, unit)
	return err
}

func (r *unitRepository) GetUnitById(id primitive.ObjectID) (unit *models.Unit, err error) {

	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&unit)
	return unit, err
}

func (r *unitRepository) GetUnitByName(name string) (unit *models.Unit, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"name": name}).Decode(&unit)
	return unit, err
}
func (r *unitRepository) GetUnitAll() ([]models.Unit, error) {
	var units []models.Unit
	result, err := r.c.Find(context.TODO(), bson.M{})
	if err != nil {
		return make([]models.Unit, 0), err
	}
	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var unit models.Unit
		if err = result.Decode(&unit); err != nil {
			return make([]models.Unit, 0), err
		}
		units = append(units, unit)

	}
	if units == nil {
		units = make([]models.Unit, 0)
	}
	return units, err
}
func (r *unitRepository) DeleteUnit(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
