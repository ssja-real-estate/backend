package main

// @title User API documentation
// @version 1.0.0
// @host localhost:8000
// @BasePath /
import (
	"log"
	// "net/http"
	"realstate/controllers"
	"realstate/db"
	_ "realstate/docs"
	"realstate/repository"
	"realstate/routes"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// func init() {
// 	err:=godotenv.Load()
// 	if err!=nil {
// 		log.Println(err)
// 	}
// }

func main() {

	con := db.NewConnection()
	defer con.Close()
	app := fiber.New()
	app.Use(logger.New())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "hello"})
	// })
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// app.Get("/swagger",func(c *fiber.Ctx) error {
	// 	return swaggerFiles.Handler
	// })

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
	usersRepo := repository.NewUsersRepository(con)
	authController := controllers.NewAuthController(usersRepo)
	authRoutes := routes.NewAuthRoutes(authController)
	authRoutes.Install(app)

	assignmenttyperepo := repository.NewAssignmentTypesRepository(con)
	assignmenttypecontoller := controllers.NewAssignmentTypeController(assignmenttyperepo)
	assignmenttyperoutes := routes.NewAssignmenttpeoute(&assignmenttypecontoller)
	assignmenttyperoutes.Install(app)

	estatetyperepo := repository.NewEstateTypesRepository(con)
	estatetypecontroller := controllers.NewEstateTypeController(estatetyperepo)
	estatetyperoutes := routes.NewEstatetypeRoute(&estatetypecontroller)
	estatetyperoutes.Install(app)

	unitrepo := repository.NewUnitRepository(con)
	unitcontroller := controllers.NewUnitController(unitrepo)
	unitroutes := routes.NewUnitRoute(&unitcontroller)
	unitroutes.Install(app)

	provincerepo := repository.NewProvinceRepository(con)
	provincecontroller := controllers.NewProvinceController(provincerepo)
	provinceroutes := routes.NewProvicneRoute(&provincecontroller)
	provinceroutes.Install(app)

	log.Fatal(app.Listen(":8000"))

}
