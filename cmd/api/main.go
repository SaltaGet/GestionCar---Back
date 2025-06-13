package main

import (
	"log"
	"os"
	"time"

	_ "github.com/DanielChachagua/GestionCar/cmd/api/docs"
	"github.com/DanielChachagua/GestionCar/cmd/api/logging"
	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"

	// "github.com/DanielChachagua/GestionCar/cmd/api/jobs"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/DanielChachagua/GestionCar/cmd/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	// "github.com/robfig/cron/v3"
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
	logging.INFO("Iniciando el servidor...")

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

	database.InitDBCache(100) 
    
  go database.StartDBJanitor()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	appDependencies := dependencies.NewApplication(db)
	// tenantDependencies := dependencies.TenantDBRepository(nil)

	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.InjectApp(appDependencies))
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

	routes.SetupRoutes(app, appDependencies)

	// repositories.Repo = dep.Repository

	app.Get("/swagger/*", swagger.HandlerDefault)

	// c := cron.New()

	// c.AddFunc("0 2 * * *", func() {
	// 	today := time.Now().Format("2006-01-02")
	// 	day := time.Now().Weekday()

	// 	if day == time.Sunday {
	// 		if err := jobs.FullBackup(today); err != nil {
	// 			log.Printf("❌ Error en backup completo: %v", err)
	// 		} else {
	// 			log.Println("✅ Backup completo exitoso.")
	// 		}
	// 	} else {
	// 		lastFull := jobs.GetLastFullBackupDir()
	// 		if err := jobs.IncrementalBackup(today, lastFull); err != nil {
	// 			log.Printf("❌ Error en backup incremental: %v", err)
	// 		} else {
	// 			log.Println("✅ Backup incremental exitoso.")
	// 		}
	// 	}
	// })
	// // c.AddFunc("0 0 1 * *", jobs.Backup)
	// // c.AddFunc("0 0 * * 0", jobs.GenerateResume)

	// c.Start()

	log.Fatal(app.Listen(":3000"))
}
