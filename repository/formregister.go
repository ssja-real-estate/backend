package repository

import (
	"realstate/db"
	"realstate/models"

	"gopkg.in/mgo.v2"
)

const formregistercollection string = "formregister"

type FormRegisterRepository interface {
	Save(formregister *models.FormRegister) error
}

type formRegisterRepository struct {
	c *mgo.Collection
}

func NewFormRegisterRepository(conn db.Connection) FormRegisterRepository {
	return &formRegisterRepository{conn.DB().C(formregistercollection)}
}

func (r *formRegisterRepository) Save(formregister *models.FormRegister) error {
	return r.c.Insert(formregister)
}
