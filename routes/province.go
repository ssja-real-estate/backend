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
	api:=app.Group("/api")
	api.Post("/province", AuthRequired, r.provinceController.CreateProvince)
	api.Put("/province", AuthRequired, r.provinceController.UpdateProvince)
	api.Get("/province/:id", AuthRequired, r.provinceController.GetProvince)
	api.Get("/province", AuthRequired, r.provinceController.GetProvinces)
	api.Delete("/province/:id", AuthRequired, r.provinceController.DeleteProvince)
	api.Post("/province/city/:id", AuthRequired, r.provinceController.AddCity)
	api.Delete("/province/city/:id", AuthRequired, r.provinceController.DeleteCity)

}
