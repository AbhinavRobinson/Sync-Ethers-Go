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
	app := fiber.New(fiber.Config{
		// Spawns child processes.
		// Prefork:               true,
		EnablePrintRoutes:     true,
		DisableStartupMessage: true,
		AppName:               "Sync Ethers API",
	})
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
	// GET Ping
	app.Get("/", func(c *fiber.Ctx) error {
		log.Trace().Msg("Ping on /")
		return c.SendString("Pong /\n")
	})
	// POST ERC20 with address
	app.Get("/erc20/:address", func(c *fiber.Ctx) error {
		address := c.Params("address")
		log.Trace().Msgf("POST /erc20/%s", address)
		status := loadToken(address)
		if status {
			tokenName, err := contracts.Tokens[address].Name(callOpts)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get token name")
				return c.SendString(fmt.Sprintf("Error Adding %s", address))
			}
			log.Trace().Msgf("Token %s added", address)
			return c.SendString(fmt.Sprintf("Added %s\n with Token: %s", c.Params("address"), tokenName))
		} else {
			log.Error().Msgf("Failed to add %s", address)
			return c.SendString(fmt.Sprintf("Error Adding %s", address))
		}
	})
	log.Info().Msg("✅ App Ready.")
}
