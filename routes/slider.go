package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type silderRoute struct {
	sliderController controllers.SliderController
}

func NewSliderRoute(sliderController controllers.SliderController) Routes {
	return &silderRoute{sliderController: sliderController}
}

func (r *silderRoute) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/carsoul", AuthRequired, r.sliderController.CreateSlider)
	api.Delete("/carsoul/:sliderId", AuthRequired, r.sliderController.DeleteSilder)
	api.Get("/carsoul", r.sliderController.GetSliders)

}
