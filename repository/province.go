package repository

import (
	"context"
	"fmt"
	"realstate/db"
	"realstate/models"
	"realstate/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const provincecollection = "province"

type ProvinceRepository interface {
	SaveProvince(province *models.Province) error
	UpdateProvince(province *models.Province, provinceid primitive.ObjectID) error
	GetProvinceById(id primitive.ObjectID) (province *models.Province, err error)
	GetProvinceByName(name string) (province *models.Province, err error)
	GetProvinceAll() (provinces []models.Province, err error)
	DeleteProvince(id primitive.ObjectID) error
	AddCity(city models.City, id primitive.ObjectID) error
	EditCity(city models.City, provinceid primitive.ObjectID, cityid primitive.ObjectID) error
	CityExists(name string, id primitive.ObjectID) (int64, error)
	GetCityById(provinceid primitive.ObjectID, cityid primitive.ObjectID) (models.City, error)
	DeleteCityByID(proviceid primitive.ObjectID, cityid primitive.ObjectID) error
	IsProvinceDelete(provinceid primitive.ObjectID) (int64, error)
	AddNeighborhood(models.Neighborhood, primitive.ObjectID, primitive.ObjectID) error
	EditNeighborhood(provinceid primitive.ObjectID, cityid primitive.ObjectID, neighborhoodid primitive.ObjectID, neighborhood models.Neighborhood) error
	GetNeighborhoodByName(provinceid primitive.ObjectID, cityid primitive.ObjectID, name string) (int64, error)
	GetNeighborhoodById(provinceid primitive.ObjectID, cityid primitive.ObjectID, neighborhoodid primitive.ObjectID) (models.Neighborhood, error)
	DeleteNeighborhoodById(provinceid primitive.ObjectID, cityid primitive.ObjectID, neghborhoodid primitive.ObjectID) error
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

func (r *provinceRepository) UpdateProvince(province *models.Province, provinceid primitive.ObjectID) error {
	filter := bson.M{"_id": provinceid}
	update := bson.M{"$set": bson.M{"name": province.Name, "mapInfo": &province.MapInfo}}
	_, err := r.c.UpdateOne(context.TODO(), filter, update)
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
		fmt.Println(err)
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
	opts := options.Update().SetUpsert(true)
	provice := bson.M{"_id": id}
	_, err := r.c.UpdateOne(context.TODO(), provice, _city, opts)
	return err

}

func (r *provinceRepository) EditCity(city models.City, provinceid primitive.ObjectID, cityid primitive.ObjectID) error {

	res := r.c.FindOneAndUpdate(
		context.TODO(),
		bson.D{
			{"_id", provinceid},
			{"cities._id", cityid},
		},
		bson.M{"$set": bson.M{"cities.$[elem]": &city}},
		options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.M{"elem._id": cityid}},
		}),
	)

	return res.Err()
}
func (r *provinceRepository) GetCityById(provinceid primitive.ObjectID, cityid primitive.ObjectID) (models.City, error) {
	var provice models.Province

	query := bson.M{"_id": provinceid, "cities._id": cityid}
	err := r.c.FindOne(context.TODO(), query).Decode(&provice)
	if err != nil {
		return models.City{}, err
	}
	for _, item := range provice.Cities {
		if item.Id == cityid {
			return item, nil
		}
	}
	return models.City{}, mongo.ErrNoDocuments
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
func (r *provinceRepository) CityExists(name string, id primitive.ObjectID) (int64, error) {

	count, err := r.c.CountDocuments(context.TODO(), bson.M{"_id": id, "cities.name": name})
	if err != nil {
		return 0, err
	}
	return count, err

}

func (r *provinceRepository) DeleteCityByID(proviceid primitive.ObjectID, cityid primitive.ObjectID) error {

	count, err := r.IsCityDelete(proviceid, cityid)
	if err != nil {
		return err
	}
	if count > 0 {
		return util.ErrNotDeleteCity
	}
	province := bson.M{"_id": proviceid}

	action := bson.M{"$pull": bson.M{"cities": bson.M{"_id": cityid}}}
	_, err = r.c.UpdateOne(context.TODO(), province, action)

	return err

}

func (r *provinceRepository) GetNeighborhoodByName(provinceid primitive.ObjectID, cityid primitive.ObjectID, name string) (int64, error) {
	count, err := r.c.CountDocuments(context.TODO(), bson.M{"_id": provinceid, "cities": bson.M{"_id": cityid}, "cities.neighborhoods": bson.M{"name": name}})
	return count, err
}

func (r *provinceRepository) EditNeighborhood(provinceid primitive.ObjectID, cityid primitive.ObjectID, neighborhoodid primitive.ObjectID, neighborhood models.Neighborhood) error {
	res := r.c.FindOneAndUpdate(
		context.TODO(),
		bson.D{
			{"_id", provinceid},
			{"cities._id", cityid},
			{"cities.neighborhoods._id", neighborhoodid},
		},
		bson.M{"$set": bson.M{"cities.$.neighborhoods.$[elem]": neighborhood}},
		options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.M{"elem._id": neighborhoodid}},
		}),
	)

	return res.Err()
}

func (r *provinceRepository) IsProvinceDelete(provinceid primitive.ObjectID) (int64, error) {

	count, err := r.c.CountDocuments(context.TODO(), bson.M{"_id": provinceid, "cities.0": bson.M{"$exists": true}})
	if err != nil {
		return 0, err
	}
	return count, err

}
func (r *provinceRepository) IsCityDelete(provinceid primitive.ObjectID, cityid primitive.ObjectID) (int64, error) {
	count, err := r.c.CountDocuments(context.TODO(),
		bson.M{"_id": provinceid, "cities._id": cityid, "cities.neghborhoods.0": bson.M{"$exists": true}})
	if err != nil {
		return 0, err
	}
	return count, err
}
func (r *provinceRepository) DeleteNeighborhoodById(provinceid primitive.ObjectID, cityid primitive.ObjectID, neghborhoodid primitive.ObjectID) error {
	province := bson.M{"_id": provinceid, "cities._id": cityid}

	action := bson.M{"$pull": bson.M{"cities.neghborhoods._id": bson.M{"_id": neghborhoodid}}}
	_, err := r.c.UpdateOne(context.TODO(), province, action)

	return err
}

func (r *provinceRepository) ErrIsNeighborhoodExists(provinceid primitive.ObjectID, cityid primitive.ObjectID, neigborhoodid primitive.ObjectID) (int64, error) {
	count, err := r.c.CountDocuments(context.TODO(), bson.M{"_id": provinceid, "cities._id": cityid, "citites.neighborhoods._id": neigborhoodid})
	if err != nil {
		return 0, err
	}
	return count, err
}

func (r *provinceRepository) GetNeighborhoodById(provinceid primitive.ObjectID, cityid primitive.ObjectID, neighborhoodid primitive.ObjectID) (models.Neighborhood, error) {
	var province models.Province
	err := r.c.FindOne(context.TODO(), bson.M{"_id": provinceid, "cities._id": cityid, "cities.neighborhoods._id": neighborhoodid}).Decode(&province)

	if err != nil {
		return models.Neighborhood{}, err
	}

	for _, item := range province.Cities {

		for _, city := range item.Neighborhoods {
			if city.Id == neighborhoodid {
				return city, nil
			}
		}

	}
	return models.Neighborhood{}, mongo.ErrNoDocuments
}
