package main

import (
	"context"
	ERC20 "sync-ethers-go/abis/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

var contracts Contracts

var ctx = context.Background()
var callOpts = &bind.CallOpts{Context: ctx, Pending: false}
var client *ethclient.Client

func loadToken(address string, contractType string) bool {
	log.Debug().Msg("Loading ABI...")

	// Dial Provider
	if client == nil {
		provider := RPC
		c, err := ethclient.Dial(provider)
		if err != nil {
			log.Fatal().Msgf("Error connecting to client: %s", err)
		}
		client = c
	}

	// Bind Token
	tokenAddress := common.HexToAddress(address)

	if contractType == "ERC20" {
		t, err := ERC20.NewToken(tokenAddress, client)
		if err != nil {
			log.Fatal().Msgf("Some error occurred in TOKEN. Err: %s", err)
		}
		log.Info().Msg("🧩 Contract processing...")
		// Add to mapping
		if contracts.Tokens == nil {
			// Allocate if not found
			contracts.Tokens = make(map[string]*ERC20.Token)
		}
		if contracts.Tokens[address] == nil {
			// Add if not found
			contracts.Tokens[address] = t
		}
	}

	// Add to DB
	addContractToDB(newContract(address, contractType))
	return true
}

func reloadTokens() {
	// Load Tokens
	log.Debug().Msg("Loading Tokens From DB...")
	getContractsFromDB()
	for _, contract := range getContractsFromDB() {
		loadToken(contract.Address, contract.Type)
	}
}
