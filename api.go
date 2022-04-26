package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/rs/zerolog/log"
)

func setupApi(host string, port int) {
	// Load Fiber
	app := fiber.New()
	// Middlewares
	app.Use(cors.New())
	app.Get("/dashboard", monitor.New())
	setupRoutes(app)
	// Listen
	err := app.Listen(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func setupRoutes(app *fiber.App) {
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		log.Trace().Msg("Ping on /")
		return c.SendString("Pong /")
	})
}
