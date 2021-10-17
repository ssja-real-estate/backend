package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type formRoute struct {
	formController controllers.FormController
}

func NewFormRoute(formcontroller controllers.FormController) Routes {
	return &formRoute{formcontroller}
}
func (r *formRoute) Install(app *fiber.App) {
	api:=app.Group("/api")
	api.Post("/form", AuthRequired, r.formController.CreateForm)
	api.Get("/form", AuthRequired, r.formController.GetForms)
	api.Get("/form/:id", AuthRequired, r.formController.GetForm)
	api.Put("/form/:id", AuthRequired, r.formController.UpdateForm)
	api.Delete("/form/:id", AuthRequired, r.formController.DeleteForm)

}
