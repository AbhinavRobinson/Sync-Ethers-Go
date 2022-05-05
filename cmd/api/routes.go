package api

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	internal "sync-ethers-go/internal"
	mongo "sync-ethers-go/internal/mongo"
	W3 "sync-ethers-go/internal/web3"
)

func ping(c *fiber.Ctx) error {
	log.Trace().Msg("Ping on /")
	return c.SendString("Pong /\n")
}

func erc20(c *fiber.Ctx) error {
	log.Trace().Msg("GET /erc20")
	return c.JSON(internal.AvailableContractTypes["ERC20"])
}

func addERC20(c *fiber.Ctx) error {
	address := c.Params("address")
	log.Trace().Msgf("PATCH /erc20/%s", address)
	status := W3.Web3.LoadToken(address, "ERC20", false)
	if status {
		tokenName, err := internal.Contracts.Tokens[address].Name(W3.Web3.CallOpts)
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
	if internal.Contracts.Tokens[address] != nil {
		mongo.DeleteContractFromDB(mongo.NewContract(address, "ERC20"))
		delete(internal.Contracts.Tokens, address)
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
	contracts := mongo.GetContractsFromDB()
	for _, contract := range contracts {
		addresses = append(addresses, common.HexToAddress(contract.Address))
	}

	// start watcher with contracts from db
	W3.StartWatcher(addresses)
	return c.SendStatus(200)
}
