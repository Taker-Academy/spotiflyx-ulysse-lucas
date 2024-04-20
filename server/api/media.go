package api

import (
	"spotiflyx/jwt"
	"spotiflyx/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MediaRoutes(app *fiber.App, db *gorm.DB, authMiddleware func(*fiber.Ctx) error){
	media := app.Group("/api/media", authMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	CreateMedia(media, db)
}

func CreateMedia(media fiber.Router, db *gorm.DB) {
	media.Post("/create", func(c *fiber.Ctx) error {
		type body struct {
			Title string `json:"title"`
			MediaType string `json:"mediaType"`
			Url string `json:"url"`
		}
		params := body{}
		if err := c.BodyParser(&params); err != nil || params.Url == "" ||
			(params.MediaType != "music" && params.MediaType != "video") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Bad request",
			})
		}

		// Create the media
		media := models.Media{}
		id, _ := jwt.GetUserID(c.Get("Authorization"), db)
		creatorID, _ := strconv.ParseUint(id, 10, 32)
		media.CreatorID = uint(creatorID)
		var err error
		if params.MediaType == "music" {
			err = CreateMusic(params.Title, params.Url, db, &media)
		} else {
			err = CreateVideo(params.Title, params.Url, db, &media)
		}
		// Check if the url is valid
		if err != nil {
			if strings.Contains(err.Error(), "Non existing id") {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"ok": false,
					"error": "invalide url",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Bad request",
			})
		}

		// Save the media to the database
		tx := db.Create(&media)
		if tx.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok": false,
				"error": "Internal server error",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{
				"mediaType": media.MediaType,
				"title": media.Title,
				"author": media.Author,
				"imgUrl": media.ImgUrl,
				"url": media.Url,
				"likes": 0,
				"favorite": false,
				"liked": false,
				"id": media.ID,
			},
		})
	})
}

func CreateMusic(title string, url string, db *gorm.DB, media *models.Media) error {
	client, ctx, err := GetSpotifyClient()
	if err != nil {
		return err
	}
	track, err := GetTrackInfo(url, ctx, client)
	if err != nil {
		return err
	}
	if (title == "") {
		media.Title = track.Name
	} else {
		media.Title = title
	}
	media.Url = string(track.URI)
	media.MediaType = "music"
	media.Author = track.Artists[0].Name
	media.ImgUrl = track.Album.Images[0].URL
	return nil
}

func CreateVideo(title string, url string, db *gorm.DB, media *models.Media) error {
	return nil;
}