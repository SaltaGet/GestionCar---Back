package main

import (
	"fmt"
	"log"
	"os"
	// "path/filepath"

	// "os/exec"
	"time"

	_ "github.com/DanielChachagua/GestionCar/cmd/api/docs"
	"github.com/DanielChachagua/GestionCar/cmd/api/logging"
	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"

	"github.com/DanielChachagua/GestionCar/cmd/api/jobs"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
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
	logging.INFO("Iniciando el servidor...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURI := os.Getenv("URI_DB")
	if dbURI == "" {
		log.Fatal("DATABASE_URI no está configurada en el archivo .env")
	}

	local := os.Getenv("LOCAL")
	if local == "true" {
		if err := jobs.GenerateSwagger(); err != nil {
			log.Fatalf("Error ejecutando swag init: %v", err)
		}
	}

	db, err := database.ConnectDB(dbURI)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer func() {
		database.CloseDB(db)
		database.CloseAllTenantDBs()
	}()

	database.InitDBCache(100)

	go database.StartDBJanitor()

	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Forwarded-For",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	appDependencies := dependencies.NewApplication(db)

	err = jobs.Migrations(appDependencies)
	if err != nil {
		log.Fatalf("Error al aplicar migraciones: %v", err)
	}

	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.InjectApp(appDependencies))
	// app.Use(middleware.AuditMiddleware())

	app.Use(limiter.New(limiter.Config{
		Max:        120,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Demasiadas peticiones. Intentá más tarde.",
			})
		},
	}))

	routes.SetupRoutes(app, appDependencies)

	app.Get("/swagger/*", swagger.HandlerDefault)

	c := cron.New()

	env := os.Getenv("ENV")
	if env == "prod" {
		_, err = c.AddFunc("0 4 * * *", func() {
			logging.INFO("⏰ [CRON] Iniciando backup diario...")
			cfg, err := jobs.LoadConfig(appDependencies)
			if err != nil {
				logging.ERROR("❌ [CRON] error leyendo config: %s", err.Error())
			}
			fmt.Println("⏰ Iniciando backup:", cfg.Databases)
			jobs.RunBackup(cfg)
			log.Println("✅ [CRON] Backup generado con éxito.")
		})
		if err != nil {
			logging.ERROR("❌ Error al crear cron job: %s", err.Error())
			panic(err)
		}
	}

	_, err = c.AddFunc("0 3 1 * *", func() {
		logging.INFO("⏰ [CRON] Iniciando resumen mensual...")
		err := jobs.GenetateResume(appDependencies)
		if err != nil {
			logging.ERROR("❌ Error al generar resumenes: %s", err.Error())
		}
		log.Println("✅ [CRON] Resumen generado con éxito.")
	})
	if err != nil {
		logging.ERROR("❌ Error al crear cron job: %s", err.Error())
		panic(err)
	}

	c.Start()

	log.Fatal(app.Listen(":3000"))
}
