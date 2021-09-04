package main

import (
	"log"
	"net/http"
	"realstate/controllers"
	"realstate/db"
	"realstate/repository"
	"realstate/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err:=godotenv.Load()
	if err!=nil {
		log.Println(err)
	}
}

func main() {
	con:=db.NewConnection()
	defer con.Close()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/",func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message":"hello"})
	})
	usersRepo := repository.NewUsersRepository(con)
	authController := controllers.NewAuthController(usersRepo)
	authRoutes := routes.NewAuthRoutes(authController)
	authRoutes.Install(app)
	log.Fatal(app.Listen(":8000"))


}