package api

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth", func(c *fiber.Ctx) error {
		return c.Next()
	})
	Signin(auth)
	Signup(auth)
}

func Signin(auth fiber.Router) {
	auth.Post("/signin", func(c *fiber.Ctx) error {
		return c.SendString("Signin")
	})
}

func Signup(auth fiber.Router) {
	auth.Post("/signup", func(c *fiber.Ctx) error {
		return c.SendString("Signup")
	})
}
