package main

import (
	"fmt"
	"os"

	"spotiflyx/api"
	"spotiflyx/db"
	"spotiflyx/jwt"
	"spotiflyx/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// connect to the database
	db, err := db.Connect()
	if err != nil {
		panic("Error when connecting to the database")
	}
	// create a new JWT auth middleware
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("SECRET_STR is not set")
	}
	authMiddleware := jwt.NewAuthMiddleware(secret)

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// add endpoints
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello World!")
		return c.SendString("Hello World!")
	})
	api.AuthRoutes(app, db)
	api.UserRoutes(app, db, authMiddleware)

	app.Listen(":3000")
}
