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
	api.Put("/estate/:estateId", AuthRequired, r.esteteContorller.UpdateEstate)
	api.Get("/estate/:estaeId", r.esteteContorller.GetEstate)
	api.Put("/estate/status/:estateId", AuthRequired, r.esteteContorller.UpdateStaus)
	api.Delete("/estate/:estateId", AuthRequired, r.esteteContorller.DeleteEstate)
	api.Get("/estate/list/user", AuthRequired, r.esteteContorller.GetEstateByUserID)
	api.Get("/estate/list/:status", AuthRequired, r.esteteContorller.GetStateByStatus)
	api.Post("/estate/search/", r.esteteContorller.SearchEstate)
	api.Get("/estates", r.esteteContorller.GetEstates)

}
