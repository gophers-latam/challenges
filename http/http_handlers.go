package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ChallengeService struct {
	ChallengeGateway
}

func (c *ChallengeService) CreateChallengeHandler(ctx *fiber.Ctx) error {
	var out Challenge

	err := ctx.BodyParser(&out)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	challenge, err := c.CreateChallenge(out)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return ctx.JSON(&challenge)
}

func (c *ChallengeService) RenderChallengeFormHandler(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}
