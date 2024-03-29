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
	GetFilterForm(form models.Form) (models.Form, error)
	GetFormForFilter(assignmenttypeid primitive.ObjectID, estatetypeid primitive.ObjectID) (form models.Form, err error)
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
	if err != nil {
		return models.Form{}, err
	}

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

func (r *formRepository) getMaxandMin(formid primitive.ObjectID, fieldid primitive.ObjectID) (int, int, error) {

	filterByFormId := bson.D{{Key: "$match", Value: bson.D{{Key: "dataForm._id", Value: formid}}}}
	flatArray := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$dataForm.sections"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	projectOut := bson.D{{Key: "$project", Value: bson.D{{Key: "items", Value: bson.D{{Key: "$filter", Value: bson.D{
		{Key: "input", Value: "$dataForm.sections.fields"},
		{Key: "as", Value: "item"},
		{Key: "cond", Value: bson.D{{Key: "$eq", Value: bson.A{"$$item._id", fieldid}}}},
	}}}}}}}

	flatArray2 := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$items"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}

	groupby := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: formid},
		{Key: "max", Value: bson.D{{Key: "$max", Value: "$items.value"}}},
		{Key: "min", Value: bson.D{{Key: "$min", Value: "$items.value"}}},
	}}}

	coursor, err := r.c.Database().Collection("estate").Aggregate(context.TODO(), mongo.Pipeline{filterByFormId, flatArray, projectOut, flatArray2, groupby})
	if err != nil {
		return 0, 0, err
	}

	defer coursor.Next(context.TODO())
	for coursor.Next(context.TODO()) {
		var estate struct {
			Id  primitive.ObjectID `bson:"_id"`
			Max int                `bson:"max"`
			Min int                `bson:"min"`
		}
		if err = coursor.Decode(&estate); err != nil {
			return 0, 0, err
		}
		return estate.Min, estate.Max, err

	}

	return 0, 0, nil

}
func (r *formRepository) GetFilterForm(form models.Form) (models.Form, error) {

	if len(form.Fields) == 0 {
		return models.Form{}, nil
	}
	for index, item := range form.Fields {
		if item.Type == 1 {
			form.Fields[index].Type = 6
		}
		// if item.Type == 8 {
		// 	listmap := make(map[string]bool)
		// 	for _, keyString := range item.Keys {
		// 		listmap[keyString] = false

		// 	}

		// 	form.Fields[index].FieldValue = listmap

		// }
	}
	return form, nil
}
func (r *formRepository) GetFormForFilter(assignmenttypeid primitive.ObjectID, estatetypeid primitive.ObjectID) (models.Form, error) {

	findForm := bson.D{{Key: "$match",
		Value: bson.D{{Key: "assignmentTypeId", Value: assignmenttypeid}, {Key: "estateTypeId", Value: estatetypeid}, {
			Key: "fields.filterable", Value: true,
		}}}}
	// flatArray := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$sections"}}}}
	// projectForm := bson.D{{Key: "$project", Value: bson.D{
	// 	{Key: "_id", Value: 1},
	// 	{Key: "title", Value: 1},
	// 	{Key: "assignmentTypeId", Value: 1},
	// 	{Key: "estateTypeId", Value: 1},
	// 	{Key: "fields", Value: bson.D{
	// 		{Key: "_id", Value: 1},
	// 		{Key: "title", Value: 1},
	// 		{Key:"type",Value: 1},
	// 		{Key: "fields", Value: bson.D{
	// 			{Key: "$filter", Value: bson.D{
	// 				{Key: "input", Value: "$sections.fields"},
	// 				{Key: "as", Value: "item"},
	// 				{Key: "cond", Value: bson.D{{Key: "$eq", Value: bson.A{"$$item.filterable", true}}}},
	// 			}},
	// 		}},
	// 	}},
	// }}}e
	// groupForm := bson.D{{Key: "$group", Value: bson.D{
	// 	{Key: "_id", Value: "$_id"},
	// 	{Key: "title", Value: bson.D{{Key: "$first", Value: "$title"}}},
	// 	{Key: "assignmentTypeId", Value: bson.D{{Key: "$first", Value: "$assignmentTypeId"}}},
	// 	{Key: "estateTypeId", Value: bson.D{{Key: "$first", Value: "$estateTypeId"}}},
	// 	{Key: "sections", Value: bson.D{{Key: "$push", Value: "$sections"}}},
	// }}}

	coursor, err := r.c.Aggregate(context.TODO(), mongo.Pipeline{findForm})

	if err != nil {
		return models.Form{}, err
	}

	defer coursor.Next(context.TODO())
	for coursor.Next(context.TODO()) {
		var form models.Form
		if err = coursor.Decode(&form); err != nil {
			return models.Form{}, err
		}

		if form.Id.IsZero() {
			return models.Form{}, mongo.ErrNoDocuments
		}
		return r.GetFilterForm(form)

	}
	return models.Form{}, mongo.ErrNoDocuments
}
