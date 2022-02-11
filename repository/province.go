package repository

import (
	"context"
	"realstate/db"
	"realstate/models"
	"realstate/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const provincecollection = "province"

type ProvinceRepository interface {
	SaveProvince(province *models.Province) error
	UpdateProvince(province *models.Province) error
	GetProvinceById(id primitive.ObjectID) (province *models.Province, err error)
	GetProvinceByName(name string) (province *models.Province, err error)
	GetProvinceAll() (provinces []models.Province, err error)
	DeleteProvince(id primitive.ObjectID) error
	AddCity(city models.City, id primitive.ObjectID) error
	GetCityByName(name string, id primitive.ObjectID) (int64, error)
	DeleteCityByID(city models.City, proviceid primitive.ObjectID) error
	IsProvinceDelete(provinceid primitive.ObjectID) (int64, error)
	AddNeighborhood(models.Neighborhood, primitive.ObjectID, primitive.ObjectID) error
	EditNeighborhood(provinceid primitive.ObjectID, cityid primitive.ObjectID, neighborhoodid primitive.ObjectID, neighborhood models.Neighborhood) error
	GetNeighborhoodByName(provinceid primitive.ObjectID, cityid primitive.ObjectID, name string) (int64, error)
}
type provinceRepository struct {
	c *mongo.Collection
}

func NewProvinceRepository(DB *mongo.Client) ProvinceRepository {
	return &provinceRepository{db.GetCollection(db.DB, provincecollection)}
}

func (r *provinceRepository) SaveProvince(province *models.Province) error {
	_, err := r.c.InsertOne(context.TODO(), province)
	return err
}

func (r *provinceRepository) UpdateProvince(province *models.Province) error {
	_, err := r.c.UpdateOne(context.TODO(), bson.M{"_id": province.Id}, province)
	return err

}

func (r *provinceRepository) GetProvinceById(id primitive.ObjectID) (province *models.Province, err error) {

	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&province)
	return province, err
}

func (r *provinceRepository) GetProvinceByName(name string) (province *models.Province, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"name": name}).Decode(&province)
	return province, err
}
func (r *provinceRepository) GetProvinceAll() ([]models.Province, error) {
	result, err := r.c.Find(context.TODO(), bson.M{})
	var provinces []models.Province
	if err != nil {
		return make([]models.Province, 0), err
	}

	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var province models.Province
		if err = result.Decode(&province); err != nil {
			return make([]models.Province, 0), err
		}
		provinces = append(provinces, province)
	}
	if provinces == nil {
		provinces = make([]models.Province, 0)
	}
	return provinces, err
}
func (r *provinceRepository) DeleteProvince(id primitive.ObjectID) error {

	count, err := r.IsProvinceDelete(id)
	if err != nil {
		return err
	}
	if count > 0 {
		return util.ErrNotDeleteProvince
	}
	_, err = r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
func (r *provinceRepository) AddCity(city models.City, id primitive.ObjectID) error {

	_city := bson.M{"$push": bson.M{"cities": city}}
	provice := bson.M{"_id": id}
	_, err := r.c.UpdateOne(context.TODO(), provice, _city)
	return err

}
func (r *provinceRepository) AddNeighborhood(neighborhood models.Neighborhood, cityid primitive.ObjectID, provinceid primitive.ObjectID) error {

	count, err := r.GetNeighborhoodByName(provinceid, cityid, neighborhood.Name)
	if err != nil {
		return err
	}
	if count > 0 {
		return util.ErrIsNeighborhoodExists
	}
	_neighborhood := bson.M{"$push": bson.M{"cities.$.neighborhoods": neighborhood}}
	provice := bson.M{"_id": provinceid, "cities._id": cityid}
	_, err = r.c.UpdateOne(context.TODO(), provice, _neighborhood)
	return err

}
func (r *provinceRepository) GetCityByName(name string, id primitive.ObjectID) (int64, error) {
	count, err := r.c.CountDocuments(context.TODO(), bson.M{"_id": id, "cities.name": name})
	return count, err

}

func (r *provinceRepository) DeleteCityByID(city models.City, proviceid primitive.ObjectID) error {
	province := bson.M{"_id": proviceid}
	action := bson.M{"$pull": bson.M{"cities": bson.M{"_id": city.Id}}}
	_, err := r.c.UpdateOne(context.TODO(), province, action)
	return err

}

func (r *provinceRepository) GetNeighborhoodByName(provinceid primitive.ObjectID, cityid primitive.ObjectID, name string) (int64, error) {
	count, err := r.c.CountDocuments(context.TODO(), bson.M{"_id": provinceid, "cities._id": cityid, "cities.neighborhoods.name": name})
	return count, err
}

func (r *provinceRepository) EditNeighborhood(provinceid primitive.ObjectID, cityid primitive.ObjectID, neighborhoodid primitive.ObjectID, neighborhood models.Neighborhood) error {
	count, err := r.GetNeighborhoodByName(provinceid, cityid, neighborhood.Name)
	if err != nil {
		return err
	}
	if count > 0 {
		return util.ErrIsNeighborhoodExists
	}
	query := bson.M{"_id": provinceid, "cities._id": cityid}
	u := []bson.M{{"$set": bson.M{"cities.$[elem]": neighborhood}},

		{"arrayFilters": bson.M{"elem.neighborhoods._id": neighborhoodid}}}

	_, err = r.c.UpdateOne(context.TODO(), query, u)
	return err
}

func (r *provinceRepository) IsProvinceDelete(provinceid primitive.ObjectID) (int64, error) {

	count, err := r.c.CountDocuments(context.TODO(), bson.M{"_id": provinceid})
	if err != nil {
		return 0, err
	}
	return count, nil

}
