package api

import (
	"github.com/gofiber/fiber/v2"
	"spotiflyx/models"
	"spotiflyx/jwt"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App, db *gorm.DB) {
	auth := app.Group("/api/auth", func(c *fiber.Ctx) error {
		return c.Next()
	})
	Signin(auth, db)
	Signup(auth, db)
}

func Signup(auth fiber.Router, db *gorm.DB) {
	auth.Post("/signup", func(c *fiber.Ctx) error {
		newUser := models.User{}
		if err := c.BodyParser(&newUser); err != nil ||
			newUser.Email == "" || newUser.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Bad request",
			})
		}

		// Check if the user already exists
		db.First(&models.User{}, "email = ?", newUser.Email).Scan(&newUser)
		if newUser.ID != 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "User already exists",
			})
		}

		// hash password
		hash, err := HashPassword(newUser.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok": false,
				"error": "Internal Server Error",
			})
		}
		newUser.Password = hash

		// Create the user
		db.Create(&newUser)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{
				"token": jwt.GetToken(string(newUser.ID)),
			},
		})
	})
}

func Signin(auth fiber.Router, db *gorm.DB) {
	auth.Post("/signin", func(c *fiber.Ctx) error {
		user := models.User{}
		if err := c.BodyParser(&user); err != nil ||
			user.Email == "" || user.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Bad request",
			})
		}

		// get user from the database
		existingUser := models.User{}
		db.First(&models.User{}, "email = ?", user.Email).Scan(&existingUser)

		// Check if the user exists and the password is correct
		if existingUser.ID == 0 || !CheckPasswordHash(user.Password, existingUser.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "email/password incorrect",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{
				"token": jwt.GetToken(string(user.ID)),
			},
		})
	})
}
