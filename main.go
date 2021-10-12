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
	app.Use(cors.New())

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

	formrepo := repository.NewFormRepositor(con)
	formcontroller := controllers.NewFormController(formrepo)
	formroute := routes.NewFormRoute(formcontroller)
	formroute.Install(app)

	log.Fatal(app.Listen(":8000"))

}
