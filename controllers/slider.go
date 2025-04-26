package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
		log.Printf("Error getting sliders: %v\n", err)
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(&sliders)
}

func (r *sliderController) CreateSlider(ctx *fiber.Ctx) error {
	var slider models.Slider
	strslider := ctx.FormValue("slider")

	if strslider == "" {
		log.Println("Error: 'slider' form value is empty")
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(fmt.Errorf("slider form value is missing")))
	}

	err := json.Unmarshal([]byte(strslider), &slider)
	if err != nil {
		log.Printf("Error unmarshalling slider JSON: %v\n", err)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	file, err := ctx.FormFile("slider")
	if err != nil {
		log.Printf("Error getting form file 'slider_image': %v\n", err)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(fmt.Errorf("could not get slider image file: %w", err)))
	}

	savePathDir := "/app/slider"
	err = os.MkdirAll(savePathDir, os.ModePerm)
	if err != nil {
		log.Printf("Error creating directory %s: %v\n", savePathDir, err)
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(fmt.Errorf("could not create upload directory: %w", err)))
	}

	savePath := filepath.Join(savePathDir, file.Filename)
	log.Printf("Attempting to save file to: %s\n", savePath)

	err = ctx.SaveFile(file, savePath)
	if err != nil {
		log.Printf("Error saving file %s: %v\n", savePath, err)
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(fmt.Errorf("could not save file: %w", err)))
	}

	log.Printf("File saved successfully to: %s\n", savePath)

	slider.Id = primitive.NewObjectID()
	slider.Path = file.Filename
	err = r.silder.CreateSlider(slider)
	if err != nil {
		log.Printf("Error creating slider in repository: %v\n", err)
		removeErr := os.Remove(savePath)
		if removeErr != nil {
			log.Printf("Error removing file %s after DB error: %v\n", savePath, removeErr)
		}
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(slider)
}

func (r *sliderController) DeleteSilder(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("sliderId"))
	if err != nil {
		log.Printf("Error parsing sliderId '%s': %v\n", ctx.Params("sliderId"), err)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrBadRole))
	}

	slider, err := r.silder.Get(id)
	if err != nil {
		log.Printf("Error getting slider with ID %s: %v\n", id.Hex(), err)
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(util.ErrNotFound))
	}

	err = r.silder.DeleteSlider(id)
	if err != nil {
		log.Printf("Error deleting slider with ID %s from repository: %v\n", id.Hex(), err)
		return ctx.Status(http.StatusInternalServerError).JSON(util.NewJError(err))
	}

	removePathDir := "/app/slider"
	removePath := filepath.Join(removePathDir, slider.Path)
	log.Printf("Attempting to remove file: %s\n", removePath)

	err = os.Remove(removePath)
	if err != nil {
		log.Printf("Error removing file %s: %v\n", removePath, err)
	} else {
		log.Printf("File removed successfully: %s\n", removePath)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Slider deleted successfully"})
}
