package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/tomiok/challenge-svc/challenges"
	"github.com/tomiok/challenge-svc/storage"
	"net/http"
)

func start(port string) error {
	app := createApp()
	return app.Listen(":" + port)
}

func createApp() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			err = ctx.Status(code).SendFile(fmt.Sprintf("./views/%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})
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
	app.Add(http.MethodGet, "/once", challengeSvc.GetChallengeByIdHandler)

	return app
}
