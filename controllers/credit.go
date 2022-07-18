package controllers

import (
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/security"
	"realstate/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreditController interface {
	Createcredit(ctx *fiber.Ctx) error
	Deletecredit(ctx *fiber.Ctx) error
}
type creditController struct {
	credit repository.CreditRepository
}

func NewCreditController(creditrepo repository.CreditRepository) CreditController {
	return &creditController{creditrepo}
}

func (r *creditController) Createcredit(ctx *fiber.Ctx) error {
	var credit models.Credit
	err := ctx.BodyParser(&credit)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	userId, err := security.GetUserByToken(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	credit.UserId = userId
	credit.Id = primitive.NewObjectID()
	credit.RegisterDate = time.Now()
	credit.RemainingDuration = credit.Duration
	err = r.credit.Save(credit)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(credit)
}
func (r *creditController) Deletecredit(ctx *fiber.Ctx) error {
	paymentID, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	err = r.credit.Delete(paymentID)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}
