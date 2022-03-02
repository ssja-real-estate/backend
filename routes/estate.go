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
	api.Get("/estate/:estaeId", AuthRequired, r.esteteContorller.GetEstate)
	api.Put("/estate/verify/:estateId", AuthRequired, r.esteteContorller.VerifiedEstate)
	api.Delete("/estate/:estateId", AuthRequired, r.esteteContorller.DeleteEstate)
	api.Get("/estate/list/user", AuthRequired, r.esteteContorller.GetEstateByUserID)
	api.Get("/estate/list/unverified", AuthRequired, r.esteteContorller.GetNotVerifiedEstate)
	api.Get("/estate/list/verified", AuthRequired, r.esteteContorller.Getverifiedestate)

}
