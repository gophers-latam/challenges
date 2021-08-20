package main

import "github.com/gofiber/fiber/v2"

func HealthCheckHandler(ctx *fiber.Ctx) error {
	_, err := ctx.WriteString("challenge service is ok")
	return err
}
