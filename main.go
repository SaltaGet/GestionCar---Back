package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DanielChachagua/GestionCar/database"
	"github.com/DanielChachagua/GestionCar/dependencies"
	_ "github.com/DanielChachagua/GestionCar/docs"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/DanielChachagua/GestionCar/repositories"
	"github.com/DanielChachagua/GestionCar/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
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
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect(os.Getenv("URI_DB"))
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer database.CloseDB(db)

	app := fiber.New()

	dep := dependencies.NewDependency(db)

	app.Use(middleware.LoggingMiddleware)
	// app.Use(middleware.AuditMiddleware())

	routes.SetupRoutes(app)

	repositories.Repo = dep.Repository

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":8080"))


}