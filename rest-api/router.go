package main

import (
	"github.com/TimiBolu/live-go-examples/rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	app := fiber.New()

	// create health check route
	app.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{"check": "Promise Card server is live!. 📦 🧧 💪🏾"})
	})

	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.TestHandler)

	return app
}
