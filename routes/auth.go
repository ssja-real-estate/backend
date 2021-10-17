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
	app.Post("api/signup", r.authController.SignUp)
	app.Post("api/signin", r.authController.SignIn)
	app.Get("/user", AuthRequired, r.authController.GetUsers)
	app.Get("/user/:id", AuthRequired, r.authController.GetUser)
	app.Put("/user/:id", AuthRequired, r.authController.PutUser)
	app.Delete("/user/:id", AuthRequired, r.authController.DeleteUser)
	app.Post("/user/verify/:mobile", r.authController.Verify)

}
