package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"})

	host := "localhost"
	port := 8080

	// Load Fiber
	app := fiber.New()
	// Middlewares
	app.Use(cors.New())
	app.Get("/dashboard", monitor.New())
	// Routes
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
