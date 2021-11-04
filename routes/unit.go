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
	api := app.Group("/api")
	api.Post("/unit", AuthRequired, r.unitController.CreateUnit)
	api.Put("/unit", AuthRequired, r.unitController.UpdateUnit)
	api.Get("/unit/:id", AuthRequired, r.unitController.GetUnit)
	api.Get("/unit", AuthRequired, r.unitController.GetUnits)
	api.Delete("/unit/:id", AuthRequired, r.unitController.DeleteUnit)

}
