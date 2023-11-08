package main

import (
	"crud-go-boilerplate-fiber/app/libraries/database"
	"crud-go-boilerplate-fiber/server"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/spf13/viper"
)

func main() {
	database.LoadConfig()
	app := fiber.New()
	config := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT",
		AllowHeaders: "Content-Type,Authorization",
	}

	app.Use(cors.New(config))
	app.Use(limiter.New(limiter.Config{
		Max:        100,             // Maximum number of requests
		Expiration: 1 * time.Minute, // Duration for which the limit counts
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	}))

	db := database.ConnectMYSQL()

	app.Get("api/v1/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": viper.GetString("app.name"),
		})
	})

	server.AppRoute(db, app)

	app.Listen(":3000")
}
