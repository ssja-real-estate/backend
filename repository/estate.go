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
	GetEstateByStatus(status int) ([]models.Estate, error)
	UpdateStatus(estaeid primitive.ObjectID, estateStatus models.EstateStatus) (int, error)
	GetEstateByUserID(userId primitive.ObjectID) ([]models.Estate, error)
	FindEstate(filterForm models.Filter) ([]models.Estate, error)
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
	estate := models.Estate{}
	err := r.c.FindOne(context.TODO(), bson.M{"_id": estateid}).Decode(&estate)
	if err != nil {
		return estate, err
	}
	return estate, err

}
func (r *estateRepository) GetEstateByStatus(status int) ([]models.Estate, error) {
	estates := []models.Estate{}
	result, err := r.c.Find(context.TODO(), bson.M{"estateStatus.status": status})
	if err != nil {
		return estates, nil
	}
	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var estate models.Estate
		if err = result.Decode(&estate); err != nil {
			return estates, nil
		}
		estates = append(estates, estate)

	}
	return estates, nil
}
func (r *estateRepository) UpdateStatus(estaeid primitive.ObjectID, estateStatus models.EstateStatus) (int, error) {
	query := bson.M{"_id": estaeid}
	update := bson.M{"$set": bson.M{"estateStatus": estateStatus}}
	res, err := r.c.UpdateOne(context.TODO(), query, update)
	return int(res.ModifiedCount), err
}

func (r *estateRepository) GetEstateByUserID(userId primitive.ObjectID) ([]models.Estate, error) {

	query := bson.M{"userId": userId}
	estates := []models.Estate{}
	result, err := r.c.Find(context.TODO(), query)

	if err != nil {
		return estates, nil
	}

	defer result.Close(context.TODO())

	for result.Next(context.TODO()) {
		var estate models.Estate
		if err = result.Decode(&estate); err != nil {
			return estates, nil
		}
		estates = append(estates, estate)
	}
	return estates, nil
}
func (r *estateRepository) FindEstate(filterForm models.Filter) ([]models.Estate, error) {
	var headFilter bson.D
	var formquery bson.D
	var estates []models.Estate
	if filterForm.Header.AssignmentTypeID.IsZero() == false {
		headFilter = append(headFilter, bson.E{Key: "dataForm.assignmentTypeId", Value: filterForm.Header.AssignmentTypeID})
	}
	if filterForm.Header.EstateTypeID.IsZero() == false {
		headFilter = append(headFilter, bson.E{Key: "dataForm.estateTypeId", Value: filterForm.Form.EstateTypeID})
	}
	if filterForm.Header.ProvinceID.IsZero() == false {
		headFilter = append(headFilter, bson.E{Key: "province._id", Value: filterForm.Header.ProvinceID})
	}
	if filterForm.Header.CityID.IsZero() == false {
		headFilter = append(headFilter, bson.E{Key: "city._id", Value: filterForm.Header.CityID})
	}
	if filterForm.Header.NeighborhoodID.IsZero() == false {
		headFilter = append(headFilter, bson.E{Key: "neighborhood._id", Value: filterForm.Header.NeighborhoodID})
	}

	for _, item := range filterForm.Form.Sections {
		for _, field := range item.Fileds {
			formquery = append(formquery, createQueryForm(field))
		}
	}

	resultquery := bson.D{{Key: "$and", Value: bson.A{headFilter, formquery}}}

	result, err := r.c.Find(context.TODO(), resultquery)

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

func createQueryForm(field models.Field) bson.E {
	var formquery bson.E

	switch field.Type {
	case 0:
		formquery = bson.E{Key: "dataForm.sections.fields.value", Value: bson.D{{Key: "$regex", Value: field.FieldValue}}}

	case 2, 3:
		formquery = bson.E{Key: "dataForm.sections.fields.value", Value: field.FieldValue}
	case 4:
		formquery = createQueryForm(field.Fields[0])

	case 6, 1:
		formquery = bson.E{Key: "dataForm.sections.fields.value", Value: bson.D{{Key: "$gte", Value: field.Min}, {Key: "$lte", Value: field.Max}}}

	}
	return formquery

}
