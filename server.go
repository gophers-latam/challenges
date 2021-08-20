package main

import "github.com/gofiber/fiber/v2"

func start(port string) {
	app := createApp()

	app.Listen(":" + port)
}

func createApp() *fiber.App {
	app := fiber.New()

	app.Add("GET", "/health", HealthCheckHandler)
	return app
}
