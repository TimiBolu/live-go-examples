package handlers

import "github.com/gofiber/fiber/v2"

func TestHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
