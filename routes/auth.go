package routes

import (
	"realstate/controllers"

	"github.com/gofiber/fiber/v2"
)

type authRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) Routes {
	return &authRoutes{authController}
}

func (r *authRoutes) Install(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/signup", r.authController.SignUp)
	api.Post("/signin", r.authController.SignIn)
	api.Get("/user", AuthRequired, r.authController.GetUsers)
	api.Get("/user/:id", AuthRequired, r.authController.GetUser)
	api.Put("/user/:id", AuthRequired, r.authController.PutUser)
	api.Delete("/user/:id", AuthRequired, r.authController.DeleteUser)
	api.Get("/verify/", r.authController.VeryfiyMobile)

}
