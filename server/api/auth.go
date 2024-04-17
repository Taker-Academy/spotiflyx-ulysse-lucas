package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App, db *gorm.DB) {
	auth := app.Group("/api/auth", func(c *fiber.Ctx) error {
		return c.Next()
	})
	Signin(auth, db)
	Signup(auth, db)
}

func Signin(auth fiber.Router, db *gorm.DB) {
	auth.Post("/signin", func(c *fiber.Ctx) error {
		return c.SendString("Signin")
	})
}

func Signup(auth fiber.Router, db *gorm.DB) {
	auth.Post("/signup", func(c *fiber.Ctx) error {
		return c.SendString("Signup")
	})
}
