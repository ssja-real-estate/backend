package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const provincecollection = "province"

type ProvinceRepository interface {
	SaveProvince(province *models.Province) error
	UpdateProvince(province *models.Province) error
	GetProvinceById(id string) (province *models.Province, err error)
	GetProvinceByName(name string) (province *models.Province, err error)
	GetProvinceAll() (provinces []*models.Province, err error)
	DeleteProvince(id string) error
	AddCity(city models.City, id string) error
	GetCityByName(name string, id string) (int, error)
	DeleteCityByID(city models.City, proviceid string) error
}

type provinceRepository struct {
	c *mgo.Collection
}

func NewProvinceRepository(conn db.Connection) ProvinceRepository {
	return &provinceRepository{conn.DB().C(provincecollection)}
}

func (r *provinceRepository) SaveProvince(province *models.Province) error {
	return r.c.Insert(province)
}
func (r *provinceRepository) UpdateProvince(province *models.Province) error {
	return r.c.UpdateId(province.Id, province)

}

func (r *provinceRepository) GetProvinceById(id string) (province *models.Province, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&province)
	return province, err
}

func (r *provinceRepository) GetProvinceByName(name string) (province *models.Province, err error) {
	err = r.c.Find(bson.M{"name": name}).One(&province)
	return province, err
}
func (r *provinceRepository) GetProvinceAll() (provinces []*models.Province, err error) {
	err = r.c.Find(bson.M{}).All(&provinces)
	return provinces, err
}
func (r *provinceRepository) DeleteProvince(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
func (r *provinceRepository) AddCity(city models.City, id string) error {

	_city := bson.M{"$push": bson.M{"cities": city}}

	provice := bson.M{"_id": bson.ObjectIdHex(id)}

	return r.c.Update(provice, _city)

}

func (r *provinceRepository) GetCityByName(name string, id string) (int, error) {

	return r.c.Find(bson.M{"cities.name": name}).Count()

}
func (r *provinceRepository) DeleteCityByID(city models.City, proviceid string) error {
	_city := bson.M{"$pull": bson.M{"cities": city}}

	provice := bson.M{"_id": bson.ObjectIdHex(proviceid)}
	return r.c.Update(provice, _city)

}
