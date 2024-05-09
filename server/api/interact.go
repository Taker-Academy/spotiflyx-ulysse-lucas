package api

import (
	"spotiflyx/models"
	"spotiflyx/jwt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InteractionsRoutes(app *fiber.App, db *gorm.DB, authMiddleware func(*fiber.Ctx) error){
	interactions := app.Group("/api/me", authMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	Like(interactions, db)
	Unlike(interactions, db)
	Favorite(interactions, db)
	Unfavorite(interactions, db)
	GetMyFavorites(interactions, db)
}

func GetMyFavorites(interactions fiber.Router, db *gorm.DB) {
	interactions.Get("/favorites", func(c *fiber.Ctx) error {
		// get the user id
		userID, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}

		// get all the favorites for the user
		var favorites []models.Interaction
		db.Model(&models.Interaction{}).Where("user_id = ? AND interaction_type = ?", userID, "favorite").Find(&favorites)

		// get the media for each favorite
		var media []models.Media
		for _, favorite := range favorites {
			var m models.Media
			db.First(&m, "id = ?", favorite.MediaID)
			media = append(media, m)
		}

		// format the output
		var output []fiber.Map
		for _, m := range media {
			output = append(output, fiber.Map{
				"mediaType": m.MediaType,
				"title": m.Title,
				"id": m.ID,
				"imgUrl": m.ImgUrl,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": output,
		})
	})
}

func Unfavorite(interactions fiber.Router, db *gorm.DB) {
	interactions.Delete("/save/:id", func(c *fiber.Ctx) error {
		// get the user id
		userID, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}

		// get the media id
		mediaID := c.Params("id")

		// check if the media exists
		media := models.Media{}
		tx := db.First(&media, "id = ?", mediaID)
		if tx.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"ok": false,
				"error": "Media not found",
			})
		}

		// check if the user has saved the media as favorite
		favorite := models.Interaction{}
		tx = db.First(&favorite, "user_id = ? AND media_id = ? AND interaction_type = ?", userID, mediaID, "favorite")
		if tx.RowsAffected == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Not saved",
			})
		}

		// delete the favorite
		db.Delete(&favorite)

		userid, _ := jwt.GetUserID(c.Get("Authorization"), db)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": GetMediaOutput(media, userid, db),
		})
	})
}

func Favorite(interactions fiber.Router, db *gorm.DB) {
	interactions.Post("/save/:id", func(c *fiber.Ctx) error {
		// get the user id
		userID, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}

		// get the media id
		mediaID := c.Params("id")

		// check if the media exists
		media := models.Media{}
		tx := db.First(&media, "id = ?", mediaID)
		if tx.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"ok": false,
				"error": "Media not found",
			})
		}

		// check if the user has already saved the media as favorite
		favorite := models.Interaction{}
		tx = db.First(&favorite, "user_id = ? AND media_id = ? AND interaction_type = ?", userID, mediaID, "favorite")
		if tx.RowsAffected != 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Already saved",
			})
		}

		// create a new favorite
		userIDInt, _ := strconv.Atoi(userID)
		mediaIDInt, _ := strconv.Atoi(mediaID)
		favorite = models.Interaction{
			UserID: uint(userIDInt),
			MediaID: uint(mediaIDInt),
			InteractionType: "favorite",
		}
		db.Create(&favorite)

		userid, _ := jwt.GetUserID(c.Get("Authorization"), db)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": GetMediaOutput(media, userid, db),
		})
	})
}

func Unlike(interactions fiber.Router, db *gorm.DB) {
	interactions.Delete("/like/:id", func(c *fiber.Ctx) error {
		// get the user id
		userID, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}

		// get the media id
		mediaID := c.Params("id")

		// check if the media exists
		media := models.Media{}
		tx := db.First(&media, "id = ?", mediaID)
		if tx.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"ok": false,
				"error": "Media not found",
			})
		}

		// check if the user has liked the media
		like := models.Interaction{}
		tx = db.First(&like, "user_id = ? AND media_id = ? AND interaction_type = ?", userID, mediaID, "like")
		if tx.RowsAffected == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Not liked",
			})
		}

		// delete the like
		db.Delete(&like)

		userid, _ := jwt.GetUserID(c.Get("Authorization"), db)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": GetMediaOutput(media, userid, db),
		})
	})
}

func Like(interactions fiber.Router, db *gorm.DB) {
	interactions.Post("/like/:id", func(c *fiber.Ctx) error {
		// get the user id
		userID, err := jwt.GetUserID(c.Get("Authorization"), db)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "Unauthorized",
			})
		}

		// get the media id
		mediaID := c.Params("id")

		// check if the media exists
		media := models.Media{}
		tx := db.First(&media, "id = ?", mediaID)
		if tx.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"ok": false,
				"error": "Media not found",
			})
		}

		// check if the user has already liked the media
		like := models.Interaction{}
		tx = db.First(&like, "user_id = ? AND media_id = ? AND interaction_type = ?", userID, mediaID, "like")
		if tx.RowsAffected != 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok": false,
				"error": "Already liked",
			})
		}

		// create a new like
		userIDInt, _ := strconv.Atoi(userID)
		mediaIDInt, _ := strconv.Atoi(mediaID)
		like = models.Interaction{
			UserID: uint(userIDInt),
			MediaID: uint(mediaIDInt),
			InteractionType: "like",
		}
		db.Create(&like)

		userid, _ := jwt.GetUserID(c.Get("Authorization"), db)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
			"data": GetMediaOutput(media, userid, db),
		})
	})
}

func GetLikes(db *gorm.DB, userID string, media string) ([]models.Interaction, bool, error) {
	// get all the likes for the media
	var likes []models.Interaction
	db.Model(&models.Interaction{}).Where("media_id = ? AND interaction_type = ?", media, "like").Find(&likes)

	// check if the user has liked the media
	var isLiked bool
	userIDInt, _ := strconv.Atoi(userID)
	for _, like := range likes {
		if like.UserID == uint(userIDInt) {
			isLiked = true
			break
		}
	}
	return likes, isLiked, nil
}

func GetFavorites(db *gorm.DB, userID string, media string) ([]models.Interaction, bool, error) {
	// get all the favorites for the media
	var favorites []models.Interaction
	tx := db.Model(&models.Interaction{}).Where("media_id = ? AND interaction_type = ?", media, "favorite").Find(&favorites)
	if tx.Error != nil {
		return nil, false, tx.Error
	}

	// check if the user has save the media as favorite
	var isFavorite bool
	userIDInt, _ := strconv.Atoi(userID)
	for _, like := range favorites {
		if like.UserID == uint(userIDInt) {
			isFavorite = true
			break
		}
	}
	return favorites, isFavorite, nil
}