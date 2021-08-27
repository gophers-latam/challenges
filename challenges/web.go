package challenges

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
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

	level = strings.ToLower(strings.TrimSpace(level))
	challengeType = strings.ToLower(strings.TrimSpace(challengeType))

	res, err := c.GetChallenges(Level(level), ChallengeType(challengeType))

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}
