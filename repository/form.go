package repository

import (
	"context"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const formcollection = "forms"

type FormRepository interface {
	SaveForm(form *models.Form) error
	GetForms() ([]models.Form, error)
	GetFormById(id primitive.ObjectID) (form models.Form, err error)
	GetForm(assignmenttypeid primitive.ObjectID, estatetypeid primitive.ObjectID) (form models.Form, err error)
	DeleteForm(id primitive.ObjectID) (err error)
	UpdateForm(id primitive.ObjectID, form *models.Form) error
	IsExitAssignmentTypeId(assignmenttypeid primitive.ObjectID) (int64, error)
	IsEstateTypeId(estatetypeid primitive.ObjectID) (int64, error)
}
type formRepository struct {
	c *mongo.Collection
}

func NewFormRepositor(DB *mongo.Client) FormRepository {
	return &formRepository{db.GetCollection(DB, formcollection)}
}

func (r *formRepository) SaveForm(form *models.Form) error {
	_, err := r.c.InsertOne(context.TODO(), &form)
	return err
}

func (r *formRepository) GetForm(assignmenttypeid primitive.ObjectID, estatetypeid primitive.ObjectID) (form models.Form, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"assignmentTypeId": assignmenttypeid, "estateTypeId": estatetypeid}).Decode(&form)
	return form, err
}
func (r *formRepository) GetForms() ([]models.Form, error) {

	result, err := r.c.Find(context.TODO(), bson.M{})
	var forms []models.Form
	if err != nil {
		return make([]models.Form, 0), err
	}

	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var form models.Form

		if err = result.Decode(&form); err != nil {
			return make([]models.Form, 0), err
		}
		forms = append(forms, form)
	}
	if forms == nil {
		forms = make([]models.Form, 0)
	}
	return forms, err
}

func (r *formRepository) GetFormById(id primitive.ObjectID) (form models.Form, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&form)
	return form, err
}

func (r *formRepository) DeleteForm(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *formRepository) UpdateForm(id primitive.ObjectID, form *models.Form) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": &form}
	_, err := r.c.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *formRepository) IsExitAssignmentTypeId(assignmenttypeid primitive.ObjectID) (int64, error) {
	result, err := r.c.CountDocuments(context.TODO(), bson.M{"assignmentTypeId": assignmenttypeid})
	if err != nil {
		return 0, err
	}

	return result, err
}

func (r *formRepository) IsEstateTypeId(estatetypeid primitive.ObjectID) (int64, error) {
	n, err := r.c.CountDocuments(context.TODO(), bson.M{"estateTypeId": estatetypeid})
	return n, err
}
