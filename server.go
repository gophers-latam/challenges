package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/gophers-latam/challenges/challenges"
	"github.com/gophers-latam/challenges/storage"
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

			// retrieve the custom status code if it's a fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			err = ctx.Status(code).SendFile(fmt.Sprintf("./views/%d.html", code))
			if err != nil {
				// in case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// return from handler
			return nil
		},
	})
	// middlewares
	app.Use(recover.New(), compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	db := storage.Get()

	challengeSvc := challenges.ChallengeService{
		ChallengeGateway: challenges.ChallengeGateway{
			DB: db,
		},
	}

	// routes
	app.Static("/", "./views")
	app.Add(http.MethodGet, "/health", HealthCheckHandler)
	app.Add(http.MethodPost, "/challenges", challengeSvc.CreateChallengeHandler)
	app.Add(http.MethodGet, "/challenges", challengeSvc.GetChallengesHandler)
	app.Add(http.MethodGet, "/once", challengeSvc.GetChallengeByIdHandler)
	app.Add(http.MethodGet, "/", challengeSvc.AddChallengeFormHandler)

	return app
}
