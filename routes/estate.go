package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type estateRoute struct {
	esteteContorller controllers.EstateController
}

func NewEstateRoute(estatecontroller controllers.EstateController) Routes {
	return &estateRoute{estatecontroller}
}
func (r *estateRoute) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/estate", AuthRequired, r.esteteContorller.CreateEstate)
	app.Get("/estate/:estaeId", AuthRequired, r.esteteContorller.GetEstate)
	app.Get("/estate/unverified", AuthRequired, r.esteteContorller.GetNotVerifiedEstate)
	app.Put("/estate/verify/:estateId", AuthRequired, r.esteteContorller.VerifiedEstate)
	app.Delete("/estate/:estateId", AuthRequired, r.esteteContorller.DeleteEstate)
	app.Get("estate/user/:userId", AuthRequired, r.esteteContorller.GetEstateByUserID)
	app.Get("/estate/all", AuthRequired, r.esteteContorller.VerifiedEstate)
}
