package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"realstate/models"
	"realstate/repository"
	"realstate/util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentController interface {
	CreateDocument(ctx *fiber.Ctx) error
	GetDocuments(ctx *fiber.Ctx) error
	DeleteDocument(ctx *fiber.Ctx) error
	GetDoucmentById(ctx *fiber.Ctx) error
}
type documentController struct {
	document repository.DocumentRepository
}

// CreateDocument implements DocumentController.
func (r *documentController) CreateDocument(ctx *fiber.Ctx) error {
	var document models.Document
	strdocument := ctx.FormValue("document")
	fmt.Println(strdocument)

	fmt.Println("----------------------------")
	err := json.Unmarshal([]byte(strdocument), &document)

	if err != nil {
		fmt.Println("1-")
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	file, err := ctx.FormFile("document")
	if err != nil {
		fmt.Println("2-")
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	err = ctx.SaveFile(file, "./documents")
	if err != nil {
		fmt.Println("3-")
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	document.Id = primitive.NewObjectID()
	document.Path = file.Filename
	err = r.document.SaveDoc(&document)
	if err != nil {
		fmt.Println("4-")
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(document)

}

// DeleteDocument implements DocumentController.
func (r *documentController) DeleteDocument(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("documentId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	document, err := r.document.GetDocumentById(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	err = r.document.DeleteDocument(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}

	err = os.RemoveAll(fmt.Sprintf("./documents/%s", document.Path))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{})
}

// GetDocuments implements DocumentController.
func (r *documentController) GetDocuments(ctx *fiber.Ctx) error {
	err, documents := r.document.GetDocumentAll()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(&documents)
}

// GetDoucmentById implements DocumentController.
func (r *documentController) GetDoucmentById(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	document, err := r.document.GetDocumentById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	return ctx.Status(http.StatusOK).JSON(document)
}

func NewDocumentController(docrepo repository.DocumentRepository) DocumentController {
	return &documentController{docrepo}
}
