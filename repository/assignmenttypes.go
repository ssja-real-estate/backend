package repository

import (
	"context"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const settingCollection = "settings"


type AssignmentTypeRepository interface {
	Save(assignmenttype *models.AssignmentType) error
	Update(assignmenttype *models.AssignmentType, assignmenttypeid primitive.ObjectID) error
	GetById(id primitive.ObjectID) (assignmenttype *models.AssignmentType, err error)
	GetByName(name string) (assignmenttype *models.AssignmentType, err error)
	GetAll() ([]models.AssignmentType, error)
	Delete(id primitive.ObjectID) error
}

type assignmentTypeRepository struct {
	c *mongo.Collection
}

func NewAssignmentTypesRepository(DB *mongo.Client) AssignmentTypeRepository {
	return &assignmentTypeRepository{db.GetCollection(DB, settingCollection)}
}
func (r *assignmentTypeRepository) Save(assignmenmenttype *models.AssignmentType) error {
	_, err := r.c.InsertOne(context.TODO(), assignmenmenttype)
	return err
}
func (r *assignmentTypeRepository) Update(assignmenttype *models.AssignmentType, assignmenttypeid primitive.ObjectID) error {
	filter := bson.M{"_id": assignmenttypeid}
	update := bson.M{"$set": &assignmenttype}
	_, err := r.c.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *assignmentTypeRepository) GetById(id primitive.ObjectID) (assignmenttype *models.AssignmentType, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&assignmenttype)
	return assignmenttype, err
}
func (r *assignmentTypeRepository) GetByName(name string) (assignmenttype *models.AssignmentType, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"name": name}).Decode(assignmenttype)
	return assignmenttype, err
}
func (r *assignmentTypeRepository) GetAll() ([]models.AssignmentType, error) {
	var assignments []models.AssignmentType
	restult, err := r.c.Find(context.TODO(), bson.M{})
	if err != nil {
		return make([]models.AssignmentType, 0), err
	}
	defer restult.Close(context.TODO())
	for restult.Next(context.TODO()) {
		var newassigment models.AssignmentType
		if err = restult.Decode(&newassigment); err != nil {
			return make([]models.AssignmentType, 0), err
		}
		assignments = append(assignments, newassigment)
	}
	if assignments == nil {
		assignments = make([]models.AssignmentType, 0)
	}
	return assignments, err
}
func (r *assignmentTypeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
