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
	app.Post("/form", AuthRequired, r.formController.CreateForm)
	app.Get("/form", AuthRequired, r.formController.GetForms)
	app.Get("/form/:id", AuthRequired, r.formController.GetForm)
	app.Put("/form/:id", AuthRequired, r.formController.UpdateForm)
	app.Delete("/form/:id", AuthRequired, r.formController.DeleteForm)

}
