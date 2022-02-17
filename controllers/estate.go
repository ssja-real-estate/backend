package controllers

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"realstate/models"
	"realstate/repository"
	"realstate/security"
	"realstate/util"
	"strings"
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
	fmt.Print(ctx)
	err := ctx.BodyParser(&estate)
	strestate := ctx.FormValue("estate")
	json.Unmarshal([]byte(strestate), &estate)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	userId, err := security.GetUserByToken(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	forms := form.File["images"]
	wd, err := os.Getwd()
	estate.Id = primitive.NewObjectID()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	images := []string{}
	for index, item := range forms {
		if index == 0 {
			err = os.Mkdir(fmt.Sprint(wd, "/app/images/", estate.Id.Hex()), fs.ModePerm)
		}
		extention := strings.Split(item.Filename, ".")[1]
		image := fmt.Sprintf("%s%d.%s", estate.Id.Hex(), index+1, extention)
		images = append(images, image)
		err = ctx.SaveFile(item, wd+"/app/images/"+estate.Id.Hex()+"/"+image)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
	}
	estate.UserId = userId
	if len(images) > 0 {
		for _, Sections := range estate.DataForm.Sections {
			for _, field := range Sections.Fileds {
				if field.Type == 5 {
					estate.DataForm.Sections[0].Fileds[0].FiledValue = images
				}
			}

		}
	}
	estate.Verified = false
	estate.CreatedAt = time.Now()
	estate.UpdateAt = time.Now()
	err = estate.DataForm.Validate()
	if err != nil {
		os.RemoveAll(estate.Id.Hex())
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	err = r.estate.SaveEstate(&estate)
	if err != nil {
		os.RemoveAll(estate.Id.Hex())
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
