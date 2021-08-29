package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tomiok/challenge-svc/challenges"
	"github.com/tomiok/challenge-svc/storage"
	"net/http"
)

func start(port string) error {
	app := createApp()
	return app.Listen(":" + port)
}

func createApp() *fiber.App {
	app := fiber.New()
	// middlewares
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	db := storage.Get()

	challengeSvc := challenges.ChallengeService{
		ChallengeGateway: challenges.ChallengeGateway{
			DB: db,
		},
	}

	// routes
	app.Add(http.MethodGet, "/health", HealthCheckHandler)
	app.Add(http.MethodPost, "/challenges", challengeSvc.CreateChallengeHandler)
	app.Add(http.MethodGet, "/challenges", challengeSvc.GetChallengesHandler)

	return app
}
