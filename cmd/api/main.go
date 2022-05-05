package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func SetupApi(host string, port int) {
	// Load Fiber
	app := fiber.New(fiber.Config{
		// Spawns child processes.
		// Prefork:               true,
		EnablePrintRoutes:     true,
		DisableStartupMessage: true,
	})
	// Middlewares
	// app.Use(cors.New())
	// app.Get("/dashboard", monitor.New())
	SetupRoutes(app)
	// Listen
	err := app.Listen(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func SetupRoutes(app *fiber.App) {
	// Routes
	// GET Ping
	app.Get("/", ping)
	// GET ERC20
	app.Get("/erc20", erc20)
	// PATCH ERC20 with address
	app.Patch("/erc20/:address", addERC20)
	// DELETE ERC20 with address
	app.Delete("/erc20/:address", deleteERC20)
	// START WATCHER
	app.Get("/start", watch)
	log.Info().Msg("✅ App Ready.")
	PostInit()
}
