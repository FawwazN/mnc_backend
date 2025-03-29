package main

import (
	"github.com/FawwazN/mnc-backend/api/config"
	"github.com/FawwazN/mnc-backend/api/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.LoadConfig()
	routers.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Welcome to MNCpay!",
		})
	})

	app.Listen(":" + config.ServerPort)
}
