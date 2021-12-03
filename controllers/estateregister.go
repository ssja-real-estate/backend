package controllers

import (
	"net/http"
	"os"
	"realstate/models"
	"realstate/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2/bson"
)

type EstateRegisterController interface {
	CreateEstateRegister(ctx *fiber.Ctx) error
}

type estateRegisterController struct {
	estateregister repository.EstateRegisterRepository
}

func NewEstateRegisterController(estateregisterrepo repository.EstateRegisterRepository) EstateRegisterController {
	return &estateRegisterController{estateregisterrepo}
}

func (r *estateRegisterController) CreateEstateRegister(ctx *fiber.Ctx) error {
	var formregister models.EstateRegister
	err := ctx.BodyParser(&formregister)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	file1, err1 := ctx.FormFile("image")
	path, err := os.Getwd()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	if err1 != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	ctx.SaveFile(file1, path)
	formregister.Id = bson.NewObjectId()
	formregister.CreatedAt = time.Now()
	formregister.UpdatedAt = time.Now()
	formregister.Accept = false

	return nil
}
