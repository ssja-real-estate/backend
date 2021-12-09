package controllers

import (
	"fmt"
	"net/http"
	"os"
	"realstate/models"
	"realstate/repository"

	"github.com/gofiber/fiber/v2"
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
	fmt.Println(file1.Filename)
	fmt.Println(file1.Size)
	path, err := os.Getwd()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	if err1 != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	fmt.Println(path)

	err = ctx.SaveFile(file1, path+"/"+file1.Filename)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": "ok"})
	// 	formregister.Id = bson.NewObjectId()
	// 	formregister.CreatedAt = time.Now()
	// 	formregister.UpdatedAt = time.Now()
	// 	formregister.Accept = false

	// 	return nil
}
