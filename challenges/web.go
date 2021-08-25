package challenges

import (
	"github.com/gofiber/fiber/v2"
	"log"
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

func (c *ChallengeService) GetChallengesHandler(ctx *fiber.Ctx) error {
	level := ctx.Query("level")
	challengeType := ctx.Query("type")

	res, err := c.GetChallenge(level, challengeType)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}
