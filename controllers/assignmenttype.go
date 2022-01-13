package controllers

import (
	"net/http"
	"realstate/db"
	"realstate/models"
	"realstate/repository"
	"realstate/util"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AssignmentTypeController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	GetAssignment(ctx *fiber.Ctx) error
	GetAssignments(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type assignmenttypeController struct {
	assignmenttyperepo repository.AssignmentTypeRepository
}

func NewAssignmentTypeController(assignmenttyperepo repository.AssignmentTypeRepository) assignmenttypeController {
	return assignmenttypeController{assignmenttyperepo}
}

// Create Assignment ... Create a new AssignmentType
// @Summary Create a new Assginmenttype
// @Description Create a new Assginmenttype
// @Tags AssignmentType
// @Success 200 {object} models.AssignmentType
// @Failure 404 {object} object
// @Router /assignmenttype/ [post]
func (c *assignmenttypeController) Create(ctx *fiber.Ctx) error {
	var assignmenttype models.AssignmentType
	err := ctx.BodyParser(&assignmenttype)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	exists, err := c.assignmenttyperepo.GetByName(assignmenttype.Name)
	if err == mgo.ErrNotFound {
		if strings.TrimSpace(assignmenttype.Name) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyName))
		}
		assignmenttype.CreatedAt = time.Now()
		assignmenttype.UpdatedAt = time.Now()
		assignmenttype.Id = bson.NewObjectId()
		err = c.assignmenttyperepo.Save(&assignmenttype)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		} else {
			return ctx.
				Status(http.StatusCreated).
				JSON(assignmenttype)
		}

	}
	if exists != nil {
		err = util.ErrNameAlreadyExists
	}
	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

// update Assignment ... update AssignmentType
// @Summary update Assginmenttype
// @Description update Assginmenttype
// @Tags AssignmentType
// @Success 200 {object} models.AssignmentType
// @Failure 404 {object} object
// @Router /assignmenttype/ [put]
func (r *assignmenttypeController) Update(ctx *fiber.Ctx) error {
	var assignmenttype models.AssignmentType
	err := ctx.BodyParser(&assignmenttype)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	if len(assignmenttype.Name) < 2 {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrEmptyName))
	}
	_, err = r.assignmenttyperepo.GetByName(assignmenttype.Name)
	if err == mgo.ErrNotFound {
		dbassignmenttype, err := r.assignmenttyperepo.GetById(assignmenttype.Id.Hex())
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrNotFound))
		}
		dbassignmenttype.UpdatedAt = time.Now()
		dbassignmenttype.Name = assignmenttype.Name
		err = r.assignmenttyperepo.Update(dbassignmenttype)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
		return ctx.Status(http.StatusOK).JSON(dbassignmenttype)
	}

	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNameAlreadyExists))

}

// GetAssignmentType ... Get AssignmentType by id
// @Summary Get AssignmentType by id
// @Description Get AssignmentType by id
// @Tags AssignmentType
// @Success 200 {object} models.AssignmentType
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /assignmenttype/id [get]
func (r *assignmenttypeController) GetAssignment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	assignmenttype, err := r.assignmenttyperepo.GetById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(assignmenttype)
}

// GetAssignmentTypes ... Get All AssignmentType
// @Summary Get All AssignmentType
// @Description Get All AssignmentType
// @Tags AssignmentType
// @Success 200 {array} models.AssignmentType
// @Failure 400 {object} object
// @Router /assignmenttypes [get]
func (r *assignmenttypeController) GetAssignments(ctx *fiber.Ctx) error {

	assignmenttypes, err := r.assignmenttyperepo.GetAll()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusOK).JSON(assignmenttypes)
}

// DeleteAssignmentType ... Delete AssignmentType by id
// @Summary Delete AssignmentType by id
// @Description Delete AssignmentType by id
// @Tags AssignmentType
// @Success 200 {object} models.AssignmentType
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /assignmenttype/id [delete]
func (r *assignmenttypeController) Delete(ctx *fiber.Ctx) error {
	var err error
	var count int
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}

	con := db.NewConnection()
	defer con.Close()
	formrepo := repository.NewFormRepositor(con)

	count, err = formrepo.IsExitAssignmentTypeId(bson.ObjectId(id))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	if count > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrNotDeleteAssignmentType)
	}
	err = r.assignmenttyperepo.Delete(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}
