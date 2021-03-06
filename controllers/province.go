package controllers

import (
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/util"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProvinceController interface {
	CreateProvince(ctx *fiber.Ctx) error
	UpdateProvince(ctx *fiber.Ctx) error
	GetProvince(ctx *fiber.Ctx) error
	GetProvinces(ctx *fiber.Ctx) error
	DeleteProvince(ctx *fiber.Ctx) error
	AddCity(ctx *fiber.Ctx) error
	EditCity(ctx *fiber.Ctx) error
	DeleteCity(ctx *fiber.Ctx) error
	AddNeighborhood(ctx *fiber.Ctx) error
	EditNeighborhood(ctx *fiber.Ctx) error
	DeleteNeighborhood(ctx *fiber.Ctx) error
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

	if err == mongo.ErrNoDocuments {
		if strings.TrimSpace(province.Name) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyName))
		}

		province.Id = primitive.NewObjectID()
		if province.Cities == nil {
			province.Cities = make([]models.City, 0)
		}
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
	provinceid, err := primitive.ObjectIDFromHex(ctx.Params("provinceId"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	var province models.Province
	err = ctx.BodyParser(&province)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	if len(province.Name) < 2 {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrEmptyName))
	}

	dbprovince, err := r.province.GetProvinceById(provinceid)

	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrNotFound))
	}
	if province.Name != dbprovince.Name {
		_, err = r.province.GetProvinceByName(province.Name)

	}

	if err == mongo.ErrNoDocuments || err == nil {

		err = r.province.UpdateProvince(&province, provinceid)
		if err != nil {

			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
		return ctx.Status(http.StatusOK).JSON(province)

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
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
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

	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	err = r.province.DeleteProvince(id)
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
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	var city models.City
	err = ctx.BodyParser(&city)
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
	count, err := r.province.CityExists(city.Name, id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	if count > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNameAlreadyExists))
	}
	city.Id = primitive.NewObjectID()
	if city.Neighborhoods == nil {
		city.Neighborhoods = make([]models.Neighborhood, 0)
	}
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

	provinceid, err := primitive.ObjectIDFromHex(ctx.Params("provinceId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrUnauthorized))
	}
	cityid, err := primitive.ObjectIDFromHex(ctx.Params("cityId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrUnauthorized))
	}
	err = r.province.DeleteCityByID(provinceid, cityid)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}

func (r *provinceController) AddNeighborhood(ctx *fiber.Ctx) error {
	proviceid, err := primitive.ObjectIDFromHex(ctx.Params("provinceId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	cityid, err := primitive.ObjectIDFromHex(ctx.Params("cityId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	var neighborhood models.Neighborhood
	err = ctx.BodyParser(&neighborhood)
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	neighborhood.Id = primitive.NewObjectID()
	err = r.province.AddNeighborhood(neighborhood, cityid, proviceid)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(neighborhood)
}

func (r *provinceController) EditCity(ctx *fiber.Ctx) error {
	proviceid, err := primitive.ObjectIDFromHex(ctx.Params("provinceId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	cityid, err := primitive.ObjectIDFromHex(ctx.Params("cityId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	var city models.City
	err = ctx.BodyParser(&city)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	dbcity, err := r.province.GetCityById(proviceid, cityid)
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	if dbcity.Name != city.Name {
		count, err := r.province.CityExists(city.Name, proviceid)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
		if count > 0 {
			return ctx.Status(http.StatusOK).JSON(dbcity)
		}
	}
	city.Id = dbcity.Id
	err = r.province.EditCity(city, proviceid, cityid)
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusOK).JSON(city)
}

func (r *provinceController) EditNeighborhood(ctx *fiber.Ctx) error {
	proviceid, err := primitive.ObjectIDFromHex(ctx.Params("provinceId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	cityid, err := primitive.ObjectIDFromHex(ctx.Params("cityId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	neighborhoodid, err := primitive.ObjectIDFromHex(ctx.Params("neighborhoodId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	dbneighborhood, err := r.province.GetNeighborhoodById(proviceid, cityid, neighborhoodid)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	var neighborhood models.Neighborhood
	err = ctx.BodyParser(&neighborhood)

	if neighborhood.Name != dbneighborhood.Name {

		count, err := r.province.GetNeighborhoodByName(proviceid, cityid, neighborhood.Name)
		if err != nil {

			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}

		if count > 0 {

			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrIsNeighborhoodExists))
		}
	}
	neighborhood.Id = dbneighborhood.Id

	err = r.province.EditNeighborhood(proviceid, cityid, neighborhoodid, neighborhood)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(neighborhood)
}

func (r *provinceController) DeleteNeighborhood(ctx *fiber.Ctx) error {
	proviceid, err := primitive.ObjectIDFromHex(ctx.Params("provinceId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	cityid, err := primitive.ObjectIDFromHex(ctx.Params("cityId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	neighborhoodId, err := primitive.ObjectIDFromHex(ctx.Params("neighborhoodId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	err = r.province.DeleteNeighborhoodById(proviceid, cityid, neighborhoodId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}
