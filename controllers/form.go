package controllers

import (
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/util"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2/bson"
)

type FormController interface {
	CreateForm(ctx *fiber.Ctx) error
}

type formController struct {
	form repository.FormRepository
}

func NewFormController(formrepo repository.FormRepository) FormController {
	return &formController{formrepo}
}

func (c *formController) CreateForm(ctx *fiber.Ctx) error {
	var form models.Form
	err := ctx.BodyParser(&form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	form.Id = bson.NewObjectId()
	err = c.form.SaveForm(&form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusCreated).JSON(form)
}
