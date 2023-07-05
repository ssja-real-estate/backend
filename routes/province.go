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
	api.Put("/province/:provinceId", AuthRequired, r.provinceController.UpdateProvince)
	api.Get("/province/:id", r.provinceController.GetProvince)
	api.Get("/province", r.provinceController.GetProvinces)
	api.Delete("/province/:id", AuthRequired, r.provinceController.DeleteProvince)
	api.Post("/province/:id/city", AuthRequired, r.provinceController.AddCity)
	api.Put("/province/:provinceId/city/:cityId", AuthRequired, r.provinceController.EditCity)
	api.Delete("/province/:provinceId/city/:cityId", AuthRequired, r.provinceController.DeleteCity)
	api.Post("/province/:provinceId/city/:cityId/neighborhood", AuthRequired, r.provinceController.AddNeighborhood)
	api.Put("/province/:provinceId/city/:cityId/neighborhood/:neighborhoodId", AuthRequired, r.provinceController.EditNeighborhood)
	api.Delete("/province/:provinceId/city/:cityId/neighborhood/:neighborhoodId", AuthRequired, r.provinceController.DeleteNeighborhood)

}
