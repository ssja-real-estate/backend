package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const settingCollection = "settings"

type AssignmentTypeRepository interface {
	Save(assignmenttype *models.AssignmentType) error
	Update(assignmenttype *models.AssignmentType) error
	GetById(id string) (assignmenttype *models.AssignmentType, err error)
	GetByName(name string) (assignmenttype *models.AssignmentType, err error)
	GetAll() (assigmenttypes []*models.AssignmentType, err error)
	Delete(id string) error
}

type assignmentTypeRepository struct {
	c *mgo.Collection
}

func NewAssignmentTypesRepository(conn db.Connection) AssignmentTypeRepository {
	return &assignmentTypeRepository{conn.DB().C(settingCollection)}
}

func (r *assignmentTypeRepository) Save(assignmenmenttype *models.AssignmentType) error {
	return r.c.Insert(assignmenmenttype)

}
func (r *assignmentTypeRepository) Update(assignmenttype *models.AssignmentType) error {
	return r.c.UpdateId(assignmenttype.Id, assignmenttype)
}

func (r *assignmentTypeRepository) GetById(id string) (assignmenttype *models.AssignmentType, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&assignmenttype)
	return assignmenttype, err
}

func (r *assignmentTypeRepository) GetByName(name string) (assignmenttype *models.AssignmentType, err error) {
	err = r.c.Find(bson.M{"name": name}).One(&assignmenttype)
	return assignmenttype, err
}
func (r *assignmentTypeRepository) GetAll() (assignments []*models.AssignmentType, err error) {

	err = r.c.Find(bson.M{}).All(&assignments)
	if assignments == nil {
		assignments = make([]*models.AssignmentType, 0)
	}
	return assignments, err
}
func (r *assignmentTypeRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
