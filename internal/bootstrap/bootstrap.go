package bootstrap

import (
	"go-starter/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Starter() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect database (optional)
	// database.InitDB()

	// Setup Route Fiber
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
