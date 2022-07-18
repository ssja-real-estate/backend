package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type paymentRoutes struct {
	paymentController controllers.PaymentController
}

func NewpaymentRoute(paymentcontroller controllers.PaymentController) Routes {
	return &paymentRoutes{paymentcontroller}
}

func (r *paymentRoutes) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/payment", AuthRequired, r.paymentController.CreatePayment)
	api.Put("/payment/:id", AuthRequired, r.paymentController.UpdatePayment)
	api.Get("/payment/:id", AuthRequired, r.paymentController.GetPaymentById)
	api.Get("/payment", AuthRequired, r.paymentController.GetPayments)
	api.Delete("/payment/:id", AuthRequired, r.paymentController.DeletePayment)
	api.Post("/paymentlink/:id", AuthRequired, r.paymentController.GetLinkPayment)

}
