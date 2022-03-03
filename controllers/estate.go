package controllers

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"realstate/db"
	"realstate/models"
	"realstate/repository"
	"realstate/security"
	"realstate/util"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstateController interface {
	CreateEstate(ctx *fiber.Ctx) error
	DeleteEstate(ctx *fiber.Ctx) error
	GetEstate(ctx *fiber.Ctx) error
	UpdateEstate(ctx *fiber.Ctx) error

	UpdateStaus(ctx *fiber.Ctx) error
	GetEstateByUserID(ctx *fiber.Ctx) error
	GetStateByStatus(ctx *fiber.Ctx) error
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
		images = append(images, fmt.Sprintf("%s/%s", estate.Id.Hex(), image))
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
					estate.DataForm.Sections[0].Fileds[0].FieldValue = images
				}
			}

		}
	}
	estate.Estatetatus.Status = 2
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
	id, err := primitive.ObjectIDFromHex(ctx.Params("estaeId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	estate, err := r.estate.GetEstateById(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(estate)
}

func (r *estateController) DeleteEstate(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("estateId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	err = r.estate.DeleteEstate(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	wd, err := os.Getwd()
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	err = os.RemoveAll(fmt.Sprintf("%s/app/images/%s", wd, id.Hex()))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{})
}

func (r *estateController) GetStateByStatus(ctx *fiber.Ctx) error {
	status, err := strconv.Atoi(ctx.Params("status"))
	estates, err := r.estate.GetEstateByStatus(status)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(estates)

}

func (r *estateController) UpdateStaus(ctx *fiber.Ctx) error {
	estaeid, err := primitive.ObjectIDFromHex(ctx.Params("estateId"))
	var EstateStatus models.EstateStatus

	ctx.BodyParser(&EstateStatus)
	userid, err := security.GetUserByToken(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	userRepo := repository.NewUsersRepository(db.DB)
	user, err := userRepo.GetById(userid)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	if user.Role != 1 {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrIsPermmisonDenied))
	}
	_, err = r.estate.UpdateStatus(estaeid, EstateStatus)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{})
}

func (r *estateController) GetEstateByUserID(ctx *fiber.Ctx) error {
	userId, err := security.GetUserByToken(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	estates, err := r.estate.GetEstateByUserID(userId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(estates)
}

func (r *estateController) UpdateEstate(ctx *fiber.Ctx) error {
	var updateestate models.Estate
	var listimages []string
	var deleteimaages []string

	estateid, err := primitive.ObjectIDFromHex(ctx.Params("estateId"))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	deletelist := ctx.FormValue("deleteimage")
	json.Unmarshal([]byte(deletelist), &deleteimaages)

	strestate := ctx.FormValue("estate")
	json.Unmarshal([]byte(strestate), &updateestate)

	oldestate, err := r.estate.GetEstateById(estateid)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))

	}

	updateestate.Id = estateid
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	forms := form.File["images"]
	wd, err := os.Getwd()

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	images := []string{}

	for _, _sections := range oldestate.DataForm.Sections {
		for _, _fileds := range _sections.Fileds {
			if _fileds.Type == 5 {
				for _, stringname := range _fileds.FieldValue.([]interface{}) {
					listimages = append(listimages, stringname.(string))
				}

			}

		}

	}

	for index, item := range forms {

		extention := strings.Split(item.Filename, ".")[1]
		image := fmt.Sprintf("%s%d.%s", updateestate.Id.Hex(), index+1, extention)
		images = append(images, fmt.Sprintf("%s/%s", updateestate.Id.Hex(), image))
		err = ctx.SaveFile(item, wd+"/app/images/"+updateestate.Id.Hex()+"/"+image)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
	}

	if len(images) > 0 {
		for _, Sections := range updateestate.DataForm.Sections {
			for _, field := range Sections.Fileds {
				if field.Type == 5 {
					updateestate.DataForm.Sections[0].Fileds[0].FieldValue = images
				}
			}

		}
	}

	updateestate.UpdateAt = time.Now()
	err = updateestate.DataForm.Validate()
	if err != nil {
		os.RemoveAll(updateestate.Id.Hex())
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	err = r.estate.SaveEstate(&updateestate)
	if err != nil {
		os.RemoveAll(updateestate.Id.Hex())
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(updateestate)

}
