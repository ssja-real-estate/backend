package controllers

import (
	"fmt"
	"net/http"
	"os"
	"realstate/models"
	"realstate/repository"
	"realstate/sadadportal"
	"realstate/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentController interface {
	CreatePayment(Ctx *fiber.Ctx) error
	UpdatePayment(Ctx *fiber.Ctx) error
	DeletePayment(Ctx *fiber.Ctx) error
	GetPaymentById(Ctx *fiber.Ctx) error
	GetPayments(Ctx *fiber.Ctx) error
	GetLinkPayment(Ctx *fiber.Ctx) error
}

type paymentContorller struct {
	payment repository.PaymentRepository
}

func NewPaymentController(paymentRepo repository.PaymentRepository) PaymentController {
	return &paymentContorller{paymentRepo}
}

// Create Payment ... Create a new Payment
// @Summary  Create New Payment
// @Description Create New Payment
// @Tags Payment
// @Success 200 {object} models.Payment
// @Failure 404 {object} object
// @Router /payment [post]
func (c *paymentContorller) CreatePayment(ctx *fiber.Ctx) error {

	var payment models.Payment
	err := ctx.BodyParser(&payment)

	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	}

	payment.Id = primitive.NewObjectID()
	err = payment.Validate()
	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	}
	isexist, err := c.payment.GetPaymentByCreditDuration(payment.Credit, payment.Duration)
	// if err = nil {
	// 	return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	// }
	if isexist {
		return ctx.Status(http.StatusBadGateway).JSON(util.ErrPaymentExists)
	}
	err = c.payment.Save(&payment)

	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(payment)
}

// update Payment ... update Payment
// @Summary update Payment
// @Description update Payment
// @Tags Payment
// @Success 200 {object} models.Payment
// @Failure 404 {object} object
// @Router /Payment/ [put]
func (c *paymentContorller) UpdatePayment(ctx *fiber.Ctx) error {
	paymentID, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	var payment models.Payment
	err = ctx.BodyParser(&payment)

	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	}
	err = payment.Validate()
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}

	err = c.payment.Update(&payment, paymentID)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(payment)

}

// DeletePayment... Delete Payment by id
// @Summary Delete Payment by id
// @Description Delete Payment by id
// @Tags Payment
// @Success 200 {object} models.Payment
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /payment/id [delete]
func (c *paymentContorller) DeletePayment(ctx *fiber.Ctx) error {
	paymentID, err := primitive.ObjectIDFromHex(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	err = c.payment.DeletePayment(paymentID)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)

}

// GetPayment ... Get Payment by id
// @Summary Get Payment by id
// @Description Get Payment by id
// @Tags Payment
// @Success 200 {object} models.Payment
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /payment/id [get]
func (c *paymentContorller) GetPaymentById(ctx *fiber.Ctx) error {

	paymentID, err := primitive.ObjectIDFromHex(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}

	var payment models.Payment
	payment, err = c.payment.GetPaymentByID(paymentID)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(payment)
}

// GetPayments ... Get All Payment
// @Summary Get All Payment
// @Description Get All Payment
// @Tags Payment
// @Success 200 {array} models.Payment
// @Failure 400 {object} object
// @Router /payment [get]
func (c *paymentContorller) GetPayments(ctx *fiber.Ctx) error {
	var payments []models.Payment
	payments, err := c.payment.GetPayments()
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(payments)

}

func (r *paymentContorller) GetLinkPayment(ctx *fiber.Ctx) error {
	paymentID, err := primitive.ObjectIDFromHex(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}

	var payment models.Payment
	payment, err = r.payment.GetPaymentByID(paymentID)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	var sadad models.SadadPayment
	sadad.Amount = payment.Credit
	sadad.LocalDateTime = time.Now()
	sadad.OrderId = 1
	sadad.MerchantId = os.Getenv("MERCHANTID")
	sadad.TerminalId = os.Getenv("TERMINALID")
	key := os.Getenv("SADADKEY")
	sadad.ReturnUrl = "url"
	singed, err := sadadportal.TripleDesECBEncrypt([]byte(fmt.Sprintf(sadad.TerminalId, ";", sadad.MerchantId, ";", sadad.Amount)), []byte(key))
	fmt.Println(err)
	sadad.SignData = singed

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"url": "https://sadad.org"})
}
