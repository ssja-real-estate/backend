package controllers

import (
	"fmt"
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/util"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProvinceController interface {
	CreateProvince(ctx *fiber.Ctx) error
	UpdateProvince(ctx *fiber.Ctx) error
	GetProvince(ctx *fiber.Ctx) error
	GetProvinces(ctx *fiber.Ctx) error
	DeleteProvince(ctx *fiber.Ctx) error
	AddCity(ctx *fiber.Ctx) error
	DeleteCity(ctx *fiber.Ctx) error
}

type provinceController struct {
	province repository.ProvinceRepository
}

func NewProvinceController(provincerepo repository.ProvinceRepository) provinceController {
	return provinceController{provincerepo}
}

// Create Province ... Create a new Province
// @Summary Create a new Province
// @Description Create a new Province
// @Tags Province
// @Success 200 {object} models.Province
// @Failure 404 {object} object
// @Router /Province [post]
func (c *provinceController) CreateProvince(ctx *fiber.Ctx) error {
	var province models.Province
	err := ctx.BodyParser(&province)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	exists, err := c.province.GetProvinceByName(province.Name)

	if err == mgo.ErrNotFound {
		if strings.TrimSpace(province.Name) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyName))
		}
		province.CreatedAt = time.Now()
		province.UpdatedAt = time.Now()
		province.Id = bson.NewObjectId()
		err = c.province.SaveProvince(&province)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		} else {
			return ctx.
				Status(http.StatusCreated).
				JSON(province)
		}

	}
	if exists != nil {
		err = util.ErrNameAlreadyExists
	}
	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

// update Province ... update Province
// @Summary update Province
// @Description update Assginmenttype
// @Tags Province
// @Success 200 {object} models.Province
// @Failure 404 {object} object
// @Router /Province/ [put]
func (r *provinceController) UpdateProvince(ctx *fiber.Ctx) error {
	var province models.Province
	err := ctx.BodyParser(&province)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	if len(province.Name) < 2 {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrEmptyName))
	}

	_, err = r.province.GetProvinceByName(province.Name)

	if err == mgo.ErrNotFound {
		dbprovince, err := r.province.GetProvinceById(province.Id.Hex())
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrNotFound))
		}
		dbprovince.UpdatedAt = time.Now()
		dbprovince.Name = province.Name
		err = r.province.UpdateProvince(dbprovince)
		if err != nil {

			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
		return ctx.Status(http.StatusOK).JSON(dbprovince)

	}

	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNameAlreadyExists))

}

// GetProvince ... Get Province by id
// @Summary Get Province by id
// @Description Get Province by id
// @Tags Province
// @Success 200 {object} models.Province
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /province/id [get]
func (r *provinceController) GetProvince(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	province, err := r.province.GetProvinceById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(province)
}

// GetProvinces ... Get All Province
// @Summary Get All Province
// @Description Get All Province
// @Tags Province
// @Success 200 {array} models.Province
// @Failure 400 {object} object
// @Router /provinces [get]
func (r *provinceController) GetProvinces(ctx *fiber.Ctx) error {

	provinces, err := r.province.GetProvinceAll()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(provinces)
}

// DeleteProvince ... Delete Province by id
// @Summary Delete Province by id
// @Description Delete Province by id
// @Tags Province
// @Success 200 {object} models.Province
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /province/id [delete]
func (r *provinceController) DeleteProvince(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	err := r.province.DeleteProvince(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}

// AddCity ...Add City to Province
// @Summary Add City to Province
// @Description Add City to Province
// @Tags Province
// @Success 200 {object} models.City
// @Param id path string true "Province ID"
// @Failure 400 {object} object
// @Router /province/city/id [post]
func (r *provinceController) AddCity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	var city models.City
	err := ctx.BodyParser(&city)
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	exists, err := r.province.GetProvinceById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	if strings.TrimSpace(exists.Name) == "" {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(util.ErrNotFound))
	}
	count, err := r.province.GetCityByName(city.Name, id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	if count > 0 {
		fmt.Println(count)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNameAlreadyExists))
	}
	city.Id = bson.NewObjectId()
	err = r.province.AddCity(city, id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(city)
}

// AddCity ...Delete City in Province
// @Summary Delete City in Province
// @Description Delete City in Province
// @Tags Province
// @Success 200 {object} models.City
// @Param id path string true "Province ID"
// @Failure 400 {object} object
// @Router /province/city/id [delete]
func (r *provinceController) DeleteCity(ctx *fiber.Ctx) error {
	var city models.City
	err := ctx.BodyParser(&city)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrUnauthorized))
	}

	err = r.province.DeleteCityByID(city, id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(city)
}
