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
	api := app.Group("/api")
	api.Post("/assignmenttype", AuthRequired, r.assignmentypecontorller.Create)
	api.Put("/assignmenttype/:assignmenttypeId", AuthRequired, r.assignmentypecontorller.Update)
	api.Get("/assignmenttype/:id", AuthRequired, r.assignmentypecontorller.GetAssignment)
	api.Get("/assignmenttype", AuthRequired, r.assignmentypecontorller.GetAssignments)
	api.Delete("/assignmenttype/:id", AuthRequired, r.assignmentypecontorller.Delete)

}
