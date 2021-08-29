package challenges

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"strings"
	"time"
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

	l := len(res)
	if l == 0 {
		return ctx.JSON(&Challenge{})
	}

	if l == 1 {
		return ctx.JSON(res[0])
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return ctx.JSON(res[r.Intn(l-1)])
}
