package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type assignmenttyperoutes struct {
	assignmentypecontorller controllers.AssignmentTypeController
}

func NewAssignmenttpeoute(assignmenttypecontroller controllers.AssignmentTypeController) Routes {
	return &assignmenttyperoutes{assignmenttypecontroller}
}

func (r *assignmenttyperoutes) Install(app *fiber.App) {
	app.Post("/assignmenttype", AuthRequired, r.assignmentypecontorller.Create)
	app.Put("/assignmenttype", AuthRequired, r.assignmentypecontorller.Update)
	app.Get("/assignmenttype/:id", AuthRequired, r.assignmentypecontorller.GetAssignment)
	app.Get("/assignmenttype", AuthRequired, r.assignmentypecontorller.GetAssignments)
	app.Delete("/assignmenttype/:id", AuthRequired, r.assignmentypecontorller.Delete)

}
