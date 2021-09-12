package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type unitRoutes struct {
	unitController controllers.UnitController
}

func NewUnitRoute(unitcontroller controllers.UnitController) Routes {
	return &unitRoutes{unitcontroller}
}

func (r *unitRoutes) Install(app *fiber.App) {
	app.Post("/unit", AuthRequired, r.unitController.CreateUnit)
	app.Put("/unit", AuthRequired, r.unitController.UpdateUnit)
	app.Get("/unit/:id", AuthRequired, r.unitController.GetUnit)
	app.Get("/units", AuthRequired, r.unitController.GetUnits)
	app.Delete("/unit/:id", AuthRequired, r.unitController.DeleteUnit)

}
