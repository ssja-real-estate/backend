package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const estateCollection = "estatetype"

type EstateTypeRepository interface {
	SaveEstateType(estatetype *models.EstateType) error
	UpdateEstateType(estatetype *models.EstateType) error
	GetEstateTypeById(id string) (estatetype *models.EstateType, err error)
	GetEstateTypeByHexId(id bson.ObjectId) (estatetype *models.EstateType, err error)
	GetEstateTypeByName(name string) (estatetype *models.EstateType, err error)
	GetEstateTypeAll() (estatetypes []*models.EstateType, err error)
	DeleteEstateType(id string) error
	
}

type estateTypeRepository struct {
	c *mgo.Collection
}

func NewEstateTypesRepository(conn db.Connection) EstateTypeRepository {
	return &estateTypeRepository{conn.DB().C(estateCollection)}
}

func (r *estateTypeRepository) SaveEstateType(estatetype *models.EstateType) error {
	return r.c.Insert(estatetype)
}
func (r *estateTypeRepository) UpdateEstateType(estatetype *models.EstateType) error {
	return r.c.UpdateId(estatetype.Id, estatetype)
}

func (r *estateTypeRepository) GetEstateTypeById(id string) (estatetype *models.EstateType, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&estatetype)
	return estatetype, err
}

func (r *estateTypeRepository) GetEstateTypeByHexId(id bson.ObjectId) (estatetype *models.EstateType, err error) {
	err = r.c.FindId(id).One(&estatetype)
	return estatetype, err
}

func (r *estateTypeRepository) GetEstateTypeByName(name string) (estatetype *models.EstateType, err error) {
	err = r.c.Find(bson.M{"name": name}).One(&estatetype)
	return estatetype, err
}
func (r *estateTypeRepository) GetEstateTypeAll() (estatetypes []*models.EstateType, err error) {
	err = r.c.Find(bson.M{}).All(&estatetypes)
	if estatetypes == nil {
		estatetypes = make([]*models.EstateType, 0)
	}
	return estatetypes, err
}
func (r *estateTypeRepository) DeleteEstateType(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
