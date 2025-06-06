package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	_ "github.com/DanielChachagua/GestionCar/cmd/api/docs"
	"github.com/DanielChachagua/GestionCar/cmd/api/jobs"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	// "github.com/DanielChachagua/GestionCar/pkg/repositories"
	"github.com/DanielChachagua/GestionCar/cmd/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

//	@title						APP GESTIONCAR
//	@version					1.0
//	@description				This is a api to app gestioncar
//	@termsOfService				http://swagger.io/terms/
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the JWT token. Example: "Bearer eyJhbGciOiJIUz..."

func main() {
	fmt.Println("Inicio app")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.ConnectDB(os.Getenv("URI_DB"))
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer func() {
		database.CloseDB(db)
		database.CloseAllTenantDBs()
	}()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	appDependencies := dependencies.NewApplication(db)
	tenantDependencies := dependencies.TenantDBRepository(nil, db)

	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.InjectApp(appDependencies, tenantDependencies))
	// app.Use(middleware.AuditMiddleware())

	app.Use(limiter.New(limiter.Config{
		Max:        120,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Demasiadas peticiones. Intentá más tarde.",
			})
		},
	}))

	routes.SetupRoutes(app, appDependencies, tenantDependencies)

	// repositories.Repo = dep.Repository

	app.Get("/swagger/*", swagger.HandlerDefault)

	c := cron.New()

	c.AddFunc("0 0 1 * *", jobs.Backup)
	c.AddFunc("0 0 * * 0", jobs.GenerateResume)

	c.Start()

	log.Fatal(app.Listen(":3000"))
}
