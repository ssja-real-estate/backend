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
	api:=app.Group("/api")
	api.Post("/estatetype", AuthRequired, r.estatetypecontroller.CreateEstateType)
	api.Put("/estatetype", AuthRequired, r.estatetypecontroller.UpdateEstateType)
	api.Get("/estatetype/:id", AuthRequired, r.estatetypecontroller.GetEstateType)
	api.Get("/estatetype", AuthRequired, r.estatetypecontroller.GetEsatteTypes)
	api.Delete("/estatetype/:id", AuthRequired, r.estatetypecontroller.DeleteEstateType)

}
