package main

// @title User API documentation
// @version 1.0.0
// @host localhost:8000
// @BasePath /
// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	db.ConnectDB()
	usersRepo := repository.NewUsersRepository(db.DB)
	authController := controllers.NewAuthController(usersRepo)
	authRoutes := routes.NewAuthRoutes(authController)
	authRoutes.Install(app)

	assignmenttyperepo := repository.NewAssignmentTypesRepository(db.DB)
	assignmenttypecontoller := controllers.NewAssignmentTypeController(assignmenttyperepo)
	assignmenttyperoutes := routes.NewAssignmenttpeoute(&assignmenttypecontoller)
	assignmenttyperoutes.Install(app)

	estatetyperepo := repository.NewEstateTypesRepository(db.DB)
	estatetypecontroller := controllers.NewEstateTypeController(estatetyperepo)
	estatetyperoutes := routes.NewEstatetypeRoute(&estatetypecontroller)
	estatetyperoutes.Install(app)

	unitrepo := repository.NewUnitRepository(db.DB)
	unitcontroller := controllers.NewUnitController(unitrepo)
	unitroutes := routes.NewUnitRoute(&unitcontroller)
	unitroutes.Install(app)

	provincerepo := repository.NewProvinceRepository(db.DB)
	provincecontroller := controllers.NewProvinceController(provincerepo)
	provinceroutes := routes.NewProvicneRoute(&provincecontroller)
	provinceroutes.Install(app)

	formrepo := repository.NewFormRepositor(db.DB)
	formcontroller := controllers.NewFormController(formrepo)
	formroute := routes.NewFormRoute(formcontroller)
	formroute.Install(app)

	estaterepo := repository.NewEstateRepository(db.DB)
	estateregistercontroller := controllers.NewEstateController(estaterepo)
	estateroute := routes.NewEstateRoute(estateregistercontroller)
	estateroute.Install(app)

	log.Fatal(app.Listen(":8000"))

}
