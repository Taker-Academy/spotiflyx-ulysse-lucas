package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB, authMiddleware func(*fiber.Ctx) error){
	user := app.Group("/api/user", authMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	user.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User")
	})
}
