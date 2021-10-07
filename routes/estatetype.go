package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type estatetyperoutes struct {
	estatetypecontroller controllers.EstateTypeController
}

func NewEstatetypeRoute(estatetypecontoller controllers.EstateTypeController) Routes {
	return &estatetyperoutes{estatetypecontoller}
}

func (r *estatetyperoutes) Install(app *fiber.App) {
	app.Post("/estatetype", AuthRequired, r.estatetypecontroller.CreateEstateType)
	app.Put("/estatetype", AuthRequired, r.estatetypecontroller.UpdateEstateType)
	app.Get("/estatetype/:id", AuthRequired, r.estatetypecontroller.GetEstateType)
	app.Get("/estatetype", AuthRequired, r.estatetypecontroller.GetEsatteTypes)
	app.Delete("/estatetype/:id", AuthRequired, r.estatetypecontroller.DeleteEstateType)

}
