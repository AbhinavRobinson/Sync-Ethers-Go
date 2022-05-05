package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func ping(c *fiber.Ctx) error {
	log.Trace().Msg("Ping on /")
	return c.SendString("Pong /\n")
}

func erc20(c *fiber.Ctx) error {
	log.Trace().Msg("GET /erc20")
	return c.JSON(AvailableContractTypes["ERC20"])
}

func addERC20(c *fiber.Ctx) error {
	address := c.Params("address")
	log.Trace().Msgf("PATCH /erc20/%s", address)
	status := web3.loadToken(address, "ERC20", false)
	if status {
		tokenName, err := contracts.Tokens[address].Name(web3.CallOpts)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get token name")
			return c.SendString(fmt.Sprintf("Error Adding %s", address))
		}
		log.Trace().Msgf("Token %s added", address)
		return c.SendString(fmt.Sprintf("Added %s\n with Token: %s\n", c.Params("address"), tokenName))
	} else {
		log.Error().Msgf("Failed to add %s", address)
		return c.SendString(fmt.Sprintf("Error Adding %s\n", address))
	}
}

func deleteERC20(c *fiber.Ctx) error {
	address := c.Params("address")
	log.Trace().Msgf("DELETE /erc20/%s", address)
	if contracts.Tokens[address] != nil {
		deleteContractFromDB(newContract(address, "ERC20"))
		delete(contracts.Tokens, address)
		log.Trace().Msgf("Token %s deleted", address)
		return c.SendString(fmt.Sprintf("Deleted %s\n", address))
	} else {
		log.Error().Msgf("Failed to delete %s", address)
		return c.SendString(fmt.Sprintf("Error Deleting %s\n", address))
	}
}

func watch(c *fiber.Ctx) error {
	log.Trace().Msg("GET /watch")
	var addresses []common.Address

	// get all addresses from db
	contracts := getContractsFromDB()
	for _, contract := range contracts {
		addresses = append(addresses, common.HexToAddress(contract.Address))
	}

	// start watcher with contracts from db
	StartWatcher(addresses)
	return c.SendStatus(200)
}
