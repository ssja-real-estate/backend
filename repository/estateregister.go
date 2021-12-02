package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
)

const estateregistercollection string = "formregister"

type EstateRegisterRepository interface {
	Save(estateregister *models.EstateRegister) error
}

type estateRegisterRepository struct {
	c *mgo.Collection
}

func NewEstateRegisterRepository(conn db.Connection) EstateRegisterRepository {
	return &estateRegisterRepository{conn.DB().C(estateregistercollection)}
}

func (r *estateRegisterRepository) Save(formregister *models.EstateRegister) error {
	return r.c.Insert(formregister)
}
