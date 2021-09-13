package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
)

const formcollection = "forms"

type FormRepository interface {
	SaveForm(form *models.Form) error
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
