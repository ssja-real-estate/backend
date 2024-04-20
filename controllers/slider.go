package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"realstate/models"
	"realstate/repository"
	"realstate/util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SliderController interface {
	CreateSlider(ctx *fiber.Ctx) error
	GetSliders(ctx *fiber.Ctx) error
	DeleteSilder(ctx *fiber.Ctx) error
}

type sliderController struct {
	silder repository.SliderRepository
}

func NewSliderController(sliderrepo repository.SliderRepository) sliderController {
	return sliderController{sliderrepo}
}
func (r *sliderController) GetSliders(ctx *fiber.Ctx) error {
	sliders, err := r.silder.GetSliders()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(&sliders)
}

func (r *sliderController) CreateSlider(ctx *fiber.Ctx) error {
	var slider models.Slider
	strslider := ctx.FormValue("slider")

	err := json.Unmarshal([]byte(strslider), &slider)

	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	file, err := ctx.FormFile("slider")
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	err = ctx.SaveFile(file, "./slider/"+file.Filename)
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	slider.Id = primitive.NewObjectID()
	slider.Path = file.Filename
	err = r.silder.CreateSlider(slider)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(slider)
}
func (r *sliderController) DeleteSilder(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("sliderId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	slider, err := r.silder.Get(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	err = r.silder.DeleteSlider(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}

	err = os.RemoveAll(fmt.Sprintf("./slider/%s", slider.Path))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{})
}
