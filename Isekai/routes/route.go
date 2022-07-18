package routes

import (
	"isekai/Isekai/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Home)

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}
