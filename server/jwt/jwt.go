package jwt

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jtoken "github.com/golang-jwt/jwt/v4"
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
