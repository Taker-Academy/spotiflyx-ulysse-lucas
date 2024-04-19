package jwt

import (
	"os"
	"time"
	"spotiflyx/models"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jtoken "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func GetSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetClaims(tokenString string) (*jtoken.Token, error) {
	token, err := jtoken.Parse(tokenString, func(token *jtoken.Token) (interface{}, error) {
		return []byte(GetSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetUserID(tokenString string, db *gorm.DB) (string, error) {
	tokenString = tokenString[7:] // Remove the Bearer prefix
	token, err := GetClaims(tokenString)
	if err != nil {
		return "", err
	}
	claims := token.Claims.(jtoken.MapClaims)

	//check if userId exist in the db
	objId, _ := claims["ID"].(string)
	user := models.User{}
	db.First(&user, "id = ?", objId)
	if user.ID == 0 {
		return "", err
	}
	return objId, nil
}

func GetToken(userID string) string {
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":    userID,
		"exp":   time.Now().Add(time.Hour * 24 * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(GetSecret()))
	if err != nil {
		return ""
	}
	return t
}

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok": false,
				"error": "wrong token",
			})
		},
	})
}
