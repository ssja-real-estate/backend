package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type estateregisterroutes struct {
	estateregistercontroller controllers.EstateRegisterController
}

func NewEstateRegisterRoute(estateregistercontroller controllers.EstateRegisterController) Routes {
	return &estateregisterroutes{estateregistercontroller}
}

func (r *estateregisterroutes) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/estate", AuthRequired, r.estateregistercontroller.CreateEstateRegister)

}
