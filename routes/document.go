package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type documentRoute struct {
	documentController controllers.DocumentController
}

func NewDocumentRoute(documentController controllers.DocumentController) Routes {
	return &documentRoute{documentController}
}

func (r *documentRoute) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/document", AuthRequired, r.documentController.CreateDocument)
	api.Get("/document/:id", AuthRequired, r.documentController.GetDoucmentById)
	api.Delete("/document/:documentId", AuthRequired, r.documentController.DeleteDocument)
	api.Get("/document", AuthRequired, r.documentController.GetDocuments)

}
