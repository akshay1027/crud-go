package main

import (
	"fiber-api/configs"
	"fiber-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // creates go-fiber instance.

	// fiber.Map is a shortcut for map[string]interface{}, useful for JSON returns.
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{
	// 		"success": "true",
	// 		"data":    "Hello from go fiber",
	// 	})
	// })

	// Run database
	configs.ConnectDB()

	// Routes handling
	routes.UserRoute(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Listen to port 5000
	app.Listen(":5000")
}
