package challenges

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(token string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deep := ctx.Request().Header.Peek("x-deep-token")
		if token != string(deep) {
			return errors.New("token is invalid or not present")
		}
		return nil
	}
}
