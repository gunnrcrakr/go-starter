package routes

import (
	"go-starter/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Welcome)
}
