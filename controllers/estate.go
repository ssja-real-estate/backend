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
	"sort"
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
	SearchEstate(ctx *fiber.Ctx) error
}

type estateController struct {
	estate repository.EstateRepository
}

func NewEstateController(estaterepo repository.EstateRepository) EstateController {
	return &estateController{estaterepo}
}
func (r *estateController) SearchEstate(ctx *fiber.Ctx) error {
	var filterForm models.Filter
	err := ctx.BodyParser(&filterForm)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	userid, err := security.GetUserByToken(ctx)
	creditrepo := repository.NewCreditRepository(db.DB)
	_, errcredit := creditrepo.GetCredit(userid)
	iscredit := false
	if errcredit != nil {
		iscredit = true
	}

	estate, err := r.estate.FindEstate(filterForm, iscredit)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusOK).JSON(estate)
}
func (r *estateController) CreateEstate(ctx *fiber.Ctx) error {

	var estate models.Estate

	// err := ctx.BodyParser(&estate)
	strestate := ctx.FormValue("estate")

	json.Unmarshal([]byte(strestate), &estate)

	// if err != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	// }
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
		image := getname(images, extention)
		images = append(images, image)
		err = ctx.SaveFile(item, wd+"/app/images/"+estate.Id.Hex()+"/"+image)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
	}
	estate.UserId = userId
	userRepo := repository.NewUsersRepository(db.DB)
	user, _ := userRepo.GetById(userId)
	estate.Phone = user.Mobile
	if len(images) > 0 {

		for index, field := range estate.DataForm.Fields {
			if field.Type == 5 {
				estate.DataForm.Fields[index].FieldValue = images

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
	images := []string{}
	estateid, err := primitive.ObjectIDFromHex(ctx.Params("estateId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	deletelist := ctx.FormValue("deletedImages")
	json.Unmarshal([]byte(deletelist), &deleteimaages)

	strestate := ctx.FormValue("estate")
	json.Unmarshal([]byte(strestate), &updateestate)
	oldestate, err := r.estate.GetEstateById(estateid)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))

	}
	updateestate.Id = estateid
	for _, fields := range oldestate.DataForm.Fields {

		if fields.Type == 5 {
			_primitive := fields.FieldValue.(primitive.A)
			bytedata, _ := json.Marshal(_primitive)
			json.Unmarshal(bytedata, &listimages)

		}

	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	forms := form.File["images"]
	wd, err := os.Getwd()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	// delete file in hard disk
	for _, deletefile := range deleteimaages {
		fmt.Println(deletefile)
		os.Remove(fmt.Sprintf("%s/app/images/%s/%s", wd, estateid.Hex(), deletefile))
	}
	// delete file name in DB
	for _, imagefilename := range listimages {
		indexfile := sort.SearchStrings(deleteimaages, imagefilename)
		if indexfile >= len(deleteimaages) {
			images = append(images, imagefilename)
		}
	}
	for _, item := range forms {
		extention := strings.Split(item.Filename, ".")[1]
		image := getname(images, extention)
		images = append(images, image)
		err = ctx.SaveFile(item, wd+"/app/images/"+updateestate.Id.Hex()+"/"+image)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
	}
	for _index, _fields := range updateestate.DataForm.Fields {

		if _fields.Type == 5 {
			updateestate.DataForm.Fields[_index].FieldValue = images
		}

	}
	updateestate.UpdateAt = time.Now()
	updateestate.Estatetatus.Status = 2

	err = updateestate.DataForm.Validate()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	err = r.estate.UpdateEstate(&updateestate, estateid)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(&updateestate)
}

func getname(images []string, extension string) string {
	var index int = len(images) + 1
	for {
		find := sort.SearchStrings(images, fmt.Sprintf("%d.%s", index, extension))
		if find == len(images) {
			return fmt.Sprintf("%d.%s", index, extension)
		} else {
			index++
		}
	}
}
