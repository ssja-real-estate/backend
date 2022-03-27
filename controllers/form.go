package controllers

import (
	"fmt"
	"net/http"
	"realstate/db"
	"realstate/models"
	"realstate/repository"
	"realstate/util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormController interface {
	CreateForm(ctx *fiber.Ctx) error
	GetForms(cts *fiber.Ctx) error
	GetFormById(ctx *fiber.Ctx) error
	GetForm(ctx *fiber.Ctx) error
	DeleteForm(ctx *fiber.Ctx) error
	UpdateForm(ctx *fiber.Ctx) error
	GetFormFilter(ctx *fiber.Ctx) error
}

type formController struct {
	form repository.FormRepository
}

func NewFormController(formrepo repository.FormRepository) FormController {
	return &formController{formrepo}
}

// Get Froms ... Get  Froms
// @Summary  Get Forms
// @Description Get Forms
// @Tags Form
// @Success 200 {array} models.Form
// @Failure 404 {object} object
// @Router /form [get]
func (r *formController) GetForms(ctx *fiber.Ctx) error {
	forms, err := r.form.GetForms()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(forms)
}

// Create From ... Create a new Froms
// @Summary  Create New Form
// @Description Create New Form
// @Tags Form
// @Success 200 {object} models.Form
// @Failure 404 {object} object
// @Router /form [post]
func (r *formController) CreateForm(ctx *fiber.Ctx) error {
	var form models.Form
	err := ctx.BodyParser(&form)

	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	}
	// to do check assignnent type and estatetype

	assignmentrepo := repository.NewAssignmentTypesRepository(db.DB)
	assignmentContoller := NewAssignmentTypeController(assignmentrepo)

	// if !form.AssignmentTypeID.Valid() {
	// 	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrAssignmentTypeIdFailed))
	// }

	_, assignmentexisterr := assignmentContoller.assignmenttyperepo.GetById(form.AssignmentTypeID)

	estaterepo := repository.NewEstateTypesRepository(db.DB)
	estateController := NewEstateTypeController(estaterepo)
	// if !form.EstateTypeID.Valid() {
	// 	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrEstatTypeIdFailed))
	// }
	_, estateerr := estateController.esstatetype.GetEstateTypeById(form.EstateTypeID)

	if assignmentexisterr != nil && estateerr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrEstateIDAssignID)
	}
	if assignmentexisterr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrAssignmentType)
	}
	if estateerr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrEstateID)
	}

	_, err = r.form.GetForm(form.AssignmentTypeID, form.EstateTypeID)
	if err == nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrFormExists)
	}

	form.Updateid()
	form.Id = primitive.NewObjectID()

	err = r.form.SaveForm(&form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusCreated).JSON(form)
}

// Get From ... Get a new Froms
// @Summary  Get Form
// @Description Get Form
// @Tags Form
// @Success 200 {object} models.Form
// @Failure 404 {object} object
// @Param id path string true "Item ID"
// @Router /form/id [get]
func (r *formController) GetFormById(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	form, err := r.form.GetFormById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	return ctx.Status(http.StatusOK).JSON(form)

}

// Get From ... Get a new Froms
// @Summary  Get Form by assignmenttypeid and estatetypeid
// @Description Get Form
// @Tags Form
// @Success 200 {object} models.Form
// @Failure 404 {object} object
// @Param assignment_type_id path string true "Item assignment_type_id"
// @Param estate_type_id path string true "Item estate_type_id"
// @Router /form/id [get]
func (r *formController) GetForm(ctx *fiber.Ctx) error {
	asignmenttypdid, err := primitive.ObjectIDFromHex(ctx.Query("assignmentTypeId"))
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.ErrAssignmentTypeIdFailed)
	}
	estateTypeId, err := primitive.ObjectIDFromHex(ctx.Query("estateTypeId"))

	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.ErrAssignmentTypeIdFailed)
	}

	form, err := r.form.GetForm(asignmenttypdid, estateTypeId)
	if err != nil {

		return ctx.Status(http.StatusOK).JSON(nil)
	}

	return ctx.Status(http.StatusOK).JSON(form)
}

// Delete From ... Delete a Form
// @Summary  Delete Form
// @Description Delete Form
// @Tags Form
// @Success 200 {object} models.Form
// @Failure 404 {object} object
// @Router /form [Delete]
func (r *formController) DeleteForm(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.ErrAssignmentTypeIdFailed)
	}
	err = r.form.DeleteForm(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}

// update From ... update a Form
// @Summary  update Form
// @Description update Form
// @Tags Form
// @Success 200 {object} models.Form
// @Failure 404 {object} object
// @Param id path string true "Item ID"
// @Router /form [put]
func (r *formController) UpdateForm(ctx *fiber.Ctx) error {
	var form models.Form
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(util.ErrAssignmentTypeIdFailed)
	}
	err = ctx.BodyParser(&form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	form.Updateid()
	err = r.form.UpdateForm(id, &form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessUpdate)
}

func (r *formController) GetFormFilter(ctx *fiber.Ctx) error {
	// assignmentTypeId, err := primitive.ObjectIDFromHex(ctx.Params("assignmentTypeId"))
	// if err != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))

	// }
	fmt.Println(ctx.Query("assignmentTypeId"))
	assignmentTypeId, err := primitive.ObjectIDFromHex(ctx.Query("assignmentTypeId"))
	fmt.Println(assignmentTypeId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	estateTypeId, err := primitive.ObjectIDFromHex(ctx.Query("estateTypeId"))

	// estateTypeId, err := primitive.ObjectIDFromHex(ctx.Params("estateTypeId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	form, err := r.form.GetFormForFilter(assignmentTypeId, estateTypeId)
	if err != nil {

		return ctx.Status(http.StatusOK).JSON(nil)
	}
	form, err = r.form.GetFilterForm(form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(form)

}
