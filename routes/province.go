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
	app.Post("/province", AuthRequired, r.provinceController.CreateProvince)
	app.Put("/province", AuthRequired, r.provinceController.UpdateProvince)
	app.Get("/province/:id", AuthRequired, r.provinceController.GetProvince)
	app.Get("/province", AuthRequired, r.provinceController.GetProvinces)
	app.Delete("/province/:id", AuthRequired, r.provinceController.DeleteProvince)
	app.Post("/province/city/:id", AuthRequired, r.provinceController.AddCity)
	app.Delete("/province/city/:id", AuthRequired, r.provinceController.DeleteCity)

}
