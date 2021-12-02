package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const formcollection = "forms"

type FormRepository interface {
	SaveForm(form *models.Form) error
	GetForms() ([]models.Form, error)
	GetFormById(id string) (form models.Form, err error)
	GetForm(assignmenttypeid bson.ObjectId, estatetypeid bson.ObjectId) (form models.Form, err error)
	DeleteForm(id string) (err error)
	UpdateForm(id string, form *models.Form) error
}
type formRepository struct {
	c *mgo.Collection
}

func NewFormRepositor(conn db.Connection) FormRepository {
	return &formRepository{conn.DB().C(formcollection)}
}
func (r *formRepository) SaveForm(form *models.Form) error {
	return r.c.Insert(form)
}

func (r *formRepository) GetForm(assignmenttypeid bson.ObjectId, estatetypeid bson.ObjectId) (form models.Form, err error) {
	err = r.c.Find(bson.M{"assignmentTypeId": assignmenttypeid, "estateTypeId": estatetypeid}).One(&form)
	return form, err
}
func (r *formRepository) GetForms() (forms []models.Form, err error) {
	err = r.c.Find(bson.M{}).All(&forms)
	if forms == nil {
		forms = make([]models.Form, 0)
	}
	return forms, err
}

func (r *formRepository) GetFormById(id string) (form models.Form, err error) {
	err = r.c.Find(bson.M{"_id": id}).One(&form)
	return form, err
}

func (r *formRepository) DeleteForm(id string) (err error) {
	return r.c.RemoveId(bson.ObjectIdHex(id))

}

func (r *formRepository) UpdateForm(id string, form *models.Form) error {
	return r.c.UpdateId(bson.ObjectIdHex(id), form)
}
