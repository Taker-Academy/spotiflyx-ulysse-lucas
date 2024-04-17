package main

import (
	"fmt"
	"os"

	"spotiflyx/api"
	"spotiflyx/db"
	//"spotiflyx/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	db, err := db.Connect()
	if err != nil {
		panic("Error when connecting to the database")
	}
	secret := os.Getenv("SECRET_STR")
	if secret == "" {
		panic("SECRET_STR is not set")
	}
	// authMiddleware := jwt.NewAuthMiddleware(secret)

	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello World!")
		return c.SendString("Hello World!")
	})
	api.AuthRoutes(app, db)

	app.Listen(":3000")
}
