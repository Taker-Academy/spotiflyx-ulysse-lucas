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

func checkEnv() {
	if os.Getenv("JWT_SECRET") == "" {
		panic("SECRET_STR is not set")
	}
	if os.Getenv("SPOTIFY_CLIENT_ID") == "" || os.Getenv("SPOTIFY_CLIENT_SECRET") == "" {
		panic("SPOTIFY_CLIENT_ID or SPOTIFY_CLIENT_SECRET is not set")
	}
	if os.Getenv("YOUTUBE_API_KEY") == "" {
		panic("YOUTUBE_API_KEY is not set")
	}
	if _, _, err := api.GetSpotifyClient(); err != nil {
		panic("Couldn't get spotify client")
	}
	fmt.Println("Spotify client connected")
	_, err := api.GetYoutubeClient()
	if err != nil {
		panic("Couldn't get youtube client")
	}
	fmt.Println("Youtube client connected")
}

func main() {
	app := fiber.New()
	// connect to the database
	db, err := db.Connect()
	if err != nil {
		panic("Error when connecting to the database")
	}

	checkEnv()
	// create a new JWT auth middleware
	secret := os.Getenv("JWT_SECRET")
	authMiddleware := jwt.NewAuthMiddleware(secret)

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Media{})
	db.AutoMigrate(&models.Interaction{})

	// add endpoints
	app.Use(cors.New())
	api.AuthRoutes(app, db)
	api.UserRoutes(app, db, authMiddleware)
	api.MediaRoutes(app, db, authMiddleware)
	api.InteractionsRoutes(app, db, authMiddleware)

	app.Listen(":3000")
}
