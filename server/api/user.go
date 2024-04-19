package api

import (
	"spotiflyx/jwt"
	"spotiflyx/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB, authMiddleware func(*fiber.Ctx) error){
	user := app.Group("/api/user", authMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	ChangePassword(user, db)
	GetUser(user, db)
	DeleteUser(user, db)
}

func DeleteUser(user fiber.Router, db *gorm.DB) {
	user.Delete("/", func(c *fiber.Ctx) error {
		// get user from the database
		id, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil || id == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "wrong jwt token",
			})
		}
		user := models.User{}
		db.First(&models.User{}, "id = ?", id).Scan(&user)

		// delete the user
		db.Delete(&user)

		// TODO: delete all the user's playlists and songs

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{},
		})
	})
}

func GetUser(user fiber.Router, db *gorm.DB) {
	user.Get("/", func(c *fiber.Ctx) error {
		// get user from the database
		id, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil || id == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "wrong jwt token",
			})
		}
		user := models.User{}
		db.First(&models.User{}, "id = ?", id).Scan(&user)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{
				"email": user.Email,
			},
		})
	})
}

func ChangePassword(user fiber.Router, db *gorm.DB) {
	user.Put("/", func(c *fiber.Ctx) error {
		type body struct {
			gorm.Model
			OldPassword string `json:"OldPassword"`
			NewPassword string `json:"NewPassword"`
		}
		user := body{}

		// parse the request body
		if err := c.BodyParser(&user); err != nil ||
			user.OldPassword == "" || user.NewPassword == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Bad request",
			})
		}

		// get user from the database
		existingUser := models.User{}
		id, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil || id == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "wrong jwt token",
			})
		}
		db.First(&models.User{}, "id = ?", id).Scan(&existingUser)

		// Check if the OldPassword is correct
		if !CheckPasswordHash(user.OldPassword, existingUser.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "old password incorrect",
			})
		}

		// hash the new password
		hash, err := HashPassword(user.NewPassword)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok": false,
				"error": "Internal Server Error",
			})
		}

		// update the user
		existingUser.Password = hash
		db.Save(&existingUser)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{},
		})
	})
}