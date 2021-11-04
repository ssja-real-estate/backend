package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const unitcollection = "units"

type UnitRepository interface {
	SaveUnit(unit *models.Unit) error
	UpdateUnit(unit *models.Unit) error
	GetUnitById(id string) (unit *models.Unit, err error)
	GetUnitByName(name string) (unit *models.Unit, err error)
	GetUnitAll() (units []*models.Unit, err error)
	DeleteUnit(id string) error
}

type unitRepository struct {
	c *mgo.Collection
}

func NewUnitRepository(conn db.Connection) UnitRepository {
	return &unitRepository{conn.DB().C(unitcollection)}
}

func (r *unitRepository) SaveUnit(unit *models.Unit) error {
	return r.c.Insert(unit)
}
func (r *unitRepository) UpdateUnit(unit *models.Unit) error {
	return r.c.UpdateId(unit.Id, unit)
}

func (r *unitRepository) GetUnitById(id string) (unit *models.Unit, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&unit)
	return unit, err
}

func (r *unitRepository) GetUnitByName(name string) (unit *models.Unit, err error) {
	err = r.c.Find(bson.M{"name": name}).One(&unit)
	return unit, err
}
func (r *unitRepository) GetUnitAll() (units []*models.Unit, err error) {
	err = r.c.Find(bson.M{}).All(&units)
	if units == nil {
		units = make([]*models.Unit, 0)
	}
	return units, err
}
func (r *unitRepository) DeleteUnit(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
