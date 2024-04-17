package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"spotiflyx/api"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello World!")
		return c.SendString("Hello World!")
	})
	api.AuthRoutes(app)

	app.Listen(":3000")
}
