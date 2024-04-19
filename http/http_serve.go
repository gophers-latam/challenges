package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"gorm.io/gorm"
)

type WebApp struct {
	*gorm.DB
	Port string
}

func (w WebApp) App() *fiber.App {
	app := w.createApp()
	return app
}

func (w WebApp) createApp() *fiber.App {
	ve := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: ve,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			err = ctx.Status(code).SendFile(fmt.Sprintf("./views/%d.html", code))
			if err != nil {
				// in case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			return nil
		},
	})

	// middlewares
	app.Use(recover.New(), compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	svc := ChallengeService{
		ChallengeGateway: ChallengeGateway{
			DB: w.DB,
		},
	}

	// routes
	app.Static("/", "./views")
	app.Add(http.MethodGet, "/", svc.RenderChallengeFormHandler)
	app.Add(http.MethodPost, "/challenges", svc.CreateChallengeHandler)
	app.Add(http.MethodGet, "/health", healthCheckHandler)

	return app
}

func healthCheckHandler(ctx *fiber.Ctx) error {
	_, err := ctx.WriteString("challenge service is ok")
	return err
}
