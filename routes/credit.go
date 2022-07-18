package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type creditRoutes struct {
	creditController controllers.CreditController
}

func NewCreditRoute(creditcontroller controllers.CreditController) Routes {
	return &creditRoutes{creditcontroller}
}

func (r *creditRoutes) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/credit", AuthRequired, r.creditController.Createcredit)

	api.Delete("/credit/:id", AuthRequired, r.creditController.Deletecredit)

}
