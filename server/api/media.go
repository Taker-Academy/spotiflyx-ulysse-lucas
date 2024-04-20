package api

import (
	"fmt"
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
	GetRecentMedia(media, db)
	GetVideoInfo(media, db)
	GetMusicInfo(media, db)
}

func GetMediaOutput(media models.Media, userID string, db *gorm.DB) fiber.Map {
	likes, liked, _ := GetLikes(db, userID, strconv.Itoa(int(media.ID)))
	_, favorite, _ := GetFavorites(db, userID, strconv.Itoa(int(media.ID)))

	return fiber.Map{
		"mediaType": media.MediaType,
		"title": media.Title,
		"author": media.Author,
		"imgUrl": media.ImgUrl,
		"url": media.Url,
		"likes": len(likes),
		"favorite": favorite,
		"liked": liked,
		"id": media.ID,
	}
}

func GetMusicInfo(media fiber.Router, db *gorm.DB) {
	media.Get("/music/:id", func(c *fiber.Ctx) error {
		// Check if the user is authenticated
		if user, _ := jwt.GetUserID(c.Get("Authorization"), db); user == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}
		// Check if the id is valid
		id := c.Params("id")
		if _, err := strconv.Atoi(id); id == "" || err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Bad request",
			})
		}
		// Check if the music exists
		media := models.Media{}
		tx := db.First(&media, "id = ?", id)
		if tx.Error != nil || media.MediaType != "music" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"ok": false,
				"error": "Music not found",
			})
		}

		userid, _ := jwt.GetUserID(c.Get("Authorization"), db)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": GetMediaOutput(media, userid, db),
		})
	})
}

func GetVideoInfo(media fiber.Router, db *gorm.DB) {
	media.Get("/video/:id", func(c *fiber.Ctx) error {
		// Check if the user is authenticated
		if user, _ := jwt.GetUserID(c.Get("Authorization"), db); user == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}
		// Check if the id is valid
		id := c.Params("id")
		if _, err := strconv.Atoi(id); id == "" || err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Bad request",
			})
		}
		// Check if the video exists
		media := models.Media{}
		tx := db.First(&media, "id = ?", id)
		if tx.Error != nil || media.MediaType != "video" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"ok": false,
				"error": "Video not found",
			})
		}

		userid, _ := jwt.GetUserID(c.Get("Authorization"), db)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": GetMediaOutput(media, userid, db),
		})
	})
}

func GetRecentMedia(media fiber.Router, db *gorm.DB) {
	media.Get("/latest", func(c *fiber.Ctx) error {
		// Check if the user is authenticated
		if user, _ := jwt.GetUserID(c.Get("Authorization"), db); user == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}
		var mediaMusicLs []models.Media
		var mediaVideoLs []models.Media
		musicLs := []fiber.Map{}
		videoLs := []fiber.Map{}

		// get the 3 most recent music
		tx := db.Where("media_type = ?", "music").Order("created_at desc").Limit(3).Find(&mediaMusicLs)
		if tx.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok": false,
				"error": "Internal server error",
			})
		}
		for _, media := range mediaMusicLs {
			musicLs = append(musicLs, fiber.Map{
				"mediaType": media.MediaType,
				"title":     media.Title,
				"id":        media.ID,
			})
		}

		// get the 3 most recent video
		tx = db.Where("media_type = ?", "video").Order("created_at desc").Limit(3).Find(&mediaVideoLs)
		if tx.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok": false,
				"error": "Internal server error",
			})
		}
		for _, media := range mediaVideoLs {
			videoLs = append(videoLs, fiber.Map{
				"mediaType": media.MediaType,
				"title":     media.Title,
				"id":        media.ID,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{
				"music": musicLs,
				"video": videoLs,
			},
		})
	})
}

func CreateMedia(media fiber.Router, db *gorm.DB) {
	media.Post("/create", func(c *fiber.Ctx) error {
		// Check if the user is authenticated
		if user, _ := jwt.GetUserID(c.Get("Authorization"), db); user == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}
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
		userid, _ := jwt.GetUserID(c.Get("Authorization"), db)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": GetMediaOutput(media, userid, db),
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
	video, err := GetYoutubeVideoInfo(url)
	if err != nil {
		return err
	}
	if (title == "") {
		media.Title = video.Items[0].Snippet.Title
	} else {
		media.Title = title
	}
	media.MediaType = "video"
	media.ImgUrl = video.Items[0].Snippet.Thumbnails.Standard.Url
	media.Author = video.Items[0].Snippet.ChannelTitle
	media.Url = fmt.Sprintf("https://www.youtube.com/embed/%s", parseVideoUrl(url))
	return nil;
}