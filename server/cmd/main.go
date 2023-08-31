package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joomcold/go-next-docker/internal/app/router"
	"github.com/joomcold/go-next-docker/internal/database"
	"github.com/joomcold/go-next-docker/internal/initializers"
)

func init() {
	initializers.Environment()
	initializers.Postgresql()
	database.Migrations()
}

func main() {
	app := fiber.New()

	// Assign routes
	router.SetupRoutes(app)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
