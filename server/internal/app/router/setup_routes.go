package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joomcold/go-next-docker/internal/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
}
