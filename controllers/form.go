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
	GetForms(cts *fiber.Ctx) error
}

type formController struct {
	form repository.FormRepository
}

func NewFormController(formrepo repository.FormRepository) FormController {
	return &formController{formrepo}
}

func (r *formController) GetForms(ctx *fiber.Ctx) error {
	forms, err := r.form.GetForms()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.NewRresult(forms))
}

func (r *formController) CreateForm(ctx *fiber.Ctx) error {
	var form models.Form
	err := ctx.BodyParser(&form)

	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	}

	form.Updateid()
	form.Id = bson.NewObjectId()

	err = r.form.SaveForm(&form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusCreated).JSON(util.NewRresult(form))
}
