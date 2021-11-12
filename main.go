package main

import (
	"repository/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.GetAll)
	app.Get("/:id?", controllers.GetOne)
	app.Post("/", controllers.Post)

	//start server on port 3008
	app.Listen(":3008")
}
