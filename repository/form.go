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
func (r *formRepository) GetForms() (forms []models.Form, err error) {
	err = r.c.Find(bson.M{}).All(&forms)
	return forms, err
}
