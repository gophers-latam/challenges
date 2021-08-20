package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomiok/challenge-svc/challenges"
	"github.com/tomiok/challenge-svc/storage"
	"net/http"
)

func start(port string) {
	app := createApp()
	app.Listen(":" + port)
}

func createApp() *fiber.App {
	app := fiber.New()
	db := storage.Get()
	challengeSvc := challenges.ChallengeService{
		ChallengeGateway: challenges.ChallengeGateway{
			DB: db,
		},
	}

	app.Add(http.MethodGet, "/health", HealthCheckHandler)
	app.Add(http.MethodPost, "/challenges", challengeSvc.CreateChallengeHandler)
	return app
}
