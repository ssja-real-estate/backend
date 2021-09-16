package controllers

import (
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

type EstateTypeController interface {
	CreateEstateType(ctx *fiber.Ctx) error
	UpdateEstateType(ctx *fiber.Ctx) error
	GetEstateType(ctx *fiber.Ctx) error
	GetEsatteTypes(ctx *fiber.Ctx) error
	DeleteEstateType(ctx *fiber.Ctx) error
}

type estatetypeController struct {
	esstatetype repository.EstateTypeRepository
}

func NewEstateTypeController(esstatetyperepo repository.EstateTypeRepository) estatetypeController {
	return estatetypeController{esstatetyperepo}
}

// Create EstateType ... Create a new EstateType
// @Summary Create a new EstateType
// @Description Create a new EstateType
// @Tags EstateType
// @body {object} models.EstateType
// @Param Body body models.EstateType true "The EstateType to create  "
// @Success 200 {object} models.EstateType
// @Failure 404 {object} object
// @Router /EstateType [post]
// @Security ApiKeyAuth
func (c *estatetypeController) CreateEstateType(ctx *fiber.Ctx) error {
	var estatetype models.EstateType
	err := ctx.BodyParser(&estatetype)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	exists, err := c.esstatetype.GetEstateTypeByName(estatetype.Name)

	if err == mgo.ErrNotFound {
		if strings.TrimSpace(estatetype.Name) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyName))
		}
		estatetype.CreatedAt = time.Now()
		estatetype.UpdatedAt = time.Now()
		estatetype.Id = bson.NewObjectId()
		err = c.esstatetype.SaveEstateType(&estatetype)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		} else {
			return ctx.
				Status(http.StatusCreated).
				JSON(estatetype)
		}

	}
	if exists != nil {
		err = util.ErrNameAlreadyExists
	}
	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

// update EstateType ... update EstateType
// @Summary update EstateType
// @Description update Assginmenttype
// @Tags EstateType
// @Success 200 {object} models.EstateType
// @Param Body body models.EstateType true "The EstateType to update  "
// @Failure 404 {object} object
// @Router /EstateType/ [put]
// @Security ApiKeyAuth
func (r *estatetypeController) UpdateEstateType(ctx *fiber.Ctx) error {

	var estatetype models.EstateType
	err := ctx.BodyParser(&estatetype)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	if len(estatetype.Name) < 2 {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrEmptyName))
	}

	_, err = r.esstatetype.GetEstateTypeByName(estatetype.Name)

	if err == mgo.ErrNotFound {
		dbestatetype, err := r.esstatetype.GetEstateTypeById(estatetype.Id.Hex())
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrNotFound))
		}
		dbestatetype.UpdatedAt = time.Now()
		dbestatetype.Name = estatetype.Name
		err = r.esstatetype.UpdateEstateType(dbestatetype)
		if err != nil {

			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
		return ctx.Status(http.StatusOK).JSON(dbestatetype)

	}

	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNameAlreadyExists))

}

// GetEstateType ... Get EstateType by id
// @Summary Get EstateType by id
// @Description Get EstateType by id
// @Tags EstateType
// @Success 200 {object} models.EstateType
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /estatetype/id [get]
// @Security ApiKeyAuth
func (r *estatetypeController) GetEstateType(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	estatetype, err := r.esstatetype.GetEstateTypeById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(estatetype)
}

// GetEstateTypes ... Get All EstateType
// @Summary Get All EstateType
// @Description Get All EstateType
// @Tags EstateType
// @Success 200 {array} models.EstateType
// @Failure 400 {object} object
// @Router /estatetypes [get]
// @Security ApiKeyAuth
func (r *estatetypeController) GetEsatteTypes(ctx *fiber.Ctx) error {

	estatetypes, err := r.esstatetype.GetEstateTypeAll()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(estatetypes)
}

// DeleteEstateType ... Delete EstateType by id
// @Summary Delete EstateType by id
// @Description Delete EstateType by id
// @Tags EstateType
// @Success 200 {object} models.EstateType
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /estatetype/id [delete]
// @Security ApiKeyAuth
func (r *estatetypeController) DeleteEstateType(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	err := r.esstatetype.DeleteEstateType(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}
