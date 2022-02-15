package controllers

import (
	"fmt"
	"net/http"

	"realstate/models"
	"realstate/repository"
	"realstate/security"
	"realstate/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstateController interface {
	CreateEstate(ctx *fiber.Ctx) error
	// UpdateEstate(ctx *fiber.Ctx) error
	DeleteEstate(ctx *fiber.Ctx) error
	GetEstate(ctx *fiber.Ctx) error
}
type estateController struct {
	estate repository.EstateRepository
}

func NewEstateController(estaterepo repository.EstateRepository) EstateController {
	return &estateController{estaterepo}
}
func (r *estateController) CreateEstate(ctx *fiber.Ctx) error {
	var estate models.Estate
	err := ctx.BodyParser(&estate)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	userId, err := security.GetUserByToken(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	form, err := ctx.MultipartForm()

	forms := form.File["image"]

	for _, item := range forms {
		fmt.Println(item.Filename)

	}

	// wd, _ := os.Getwd()
	// file1, err1 := ctx.FormFile("image")

	// if err1 != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	// }

	// err = ctx.SaveFile(file1, wd+"/app/images/"+file1.Filename)
	// if err != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	// }
	estate.Id = primitive.NewObjectID()
	estate.UserId = userId
	estate.Verified = false
	estate.CreatedAt = time.Now()
	estate.UpdateAt = time.Now()
	err = estate.DataForm.Validate()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	err = r.estate.SaveEstate(&estate)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(estate)
}

func (r *estateController) GetEstate(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	estate, err := r.estate.GetEstateById(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": &estate})
}
func (r *estateController) DeleteEstate(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	err = r.estate.DeleteEstate(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": "The Estate id Deleted....."})
}
