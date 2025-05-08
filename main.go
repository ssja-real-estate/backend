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

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{BodyLimit: 5 * 1024 * 1024})
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://ssja.ir https://ssja.ir", // بین دامنه‌ها فاصله است نه کاما
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
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

	paymentrepo := repository.NewPaymentRepository(db.DB)
	paymentController := controllers.NewPaymentController(paymentrepo)
	paymentroute := routes.NewpaymentRoute(paymentController)
	paymentroute.Install(app)

	creditrepo := repository.NewCreditRepository(db.DB)
	creditController := controllers.NewCreditController(creditrepo)
	creditroute := routes.NewCreditRoute(creditController)
	creditroute.Install(app)

	documentrepo := repository.NewDocumentRepository(db.DB)
	documentController := controllers.NewDocumentController(documentrepo)
	documentRoute := routes.NewDocumentRoute(documentController)
	documentRoute.Install(app)

	sliderrepo := repository.NewSliderRepository(db.DB)
	sliderController := controllers.NewSliderController(sliderrepo)
	sliderRoute := routes.NewSliderRoute(&sliderController)
	sliderRoute.Install(app)

	log.Fatal(app.Listen(":8000"))

}
