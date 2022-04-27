package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func ping(c *fiber.Ctx) error {
	log.Trace().Msg("Ping on /")
	return c.SendString("Pong /\n")
}

func addERC20(c *fiber.Ctx) error {
	address := c.Params("address")
	log.Trace().Msgf("PATCH /erc20/%s", address)
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
}
