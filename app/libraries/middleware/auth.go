package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Auth interface {
	CheckAuth(ctx *fiber.Ctx) error
}
type auth struct{}

func NewAuth() Auth {
	return &auth{}
}

var secretKey = viper.GetString("app.secret")

func (a *auth) CheckAuth(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	claims := jwt.StandardClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !parsedToken.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Mengambil UserID dari token
	userid, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Menyimpan UserID dalam context Fiber
	c.Locals("userid", userid)
	return c.Next()
}
