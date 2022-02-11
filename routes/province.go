package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type provinceRoutes struct {
	provinceController controllers.ProvinceController
}

func NewProvicneRoute(provincecontroller controllers.ProvinceController) Routes {
	return &provinceRoutes{provincecontroller}
}

func (r *provinceRoutes) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/province", AuthRequired, r.provinceController.CreateProvince)
	api.Put("/province", AuthRequired, r.provinceController.UpdateProvince)
	api.Get("/province/:id", AuthRequired, r.provinceController.GetProvince)
	api.Get("/province", AuthRequired, r.provinceController.GetProvinces)
	api.Delete("/province/:id", AuthRequired, r.provinceController.DeleteProvince)
	api.Post("/province/:id/city", AuthRequired, r.provinceController.AddCity)
	api.Put("/province/:id/city", AuthRequired, r.provinceController.AddCity)
	api.Delete("/province/city/:id", AuthRequired, r.provinceController.DeleteCity)
	api.Post("/province/:provinceId/city/:cityId/neighborhood", AuthRequired, r.provinceController.AddNeighborhood)
	api.Put("/province/:provinceId/city/:cityId/neighborhood/:neighborhoodId", AuthRequired, r.provinceController.EditNeighborhood)
	api.Delete("/province/:province-id/city/city-id/neighborhood/neighborhood-id", AuthRequired, r.provinceController.DeleteNeighborhood)

}
