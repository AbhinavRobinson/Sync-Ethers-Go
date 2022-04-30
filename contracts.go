package main

import (
	ERC20 "sync-ethers-go/abis/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

var contracts Contracts

func (web3 *Web3Client) loadToken(address string, contractType string, preload bool) bool {
	log.Debug().Msg("Loading ABI...")

	// Dial Provider
	if web3 == nil {
		provider := RPC
		c, err := ethclient.Dial(provider)
		if err != nil {
			log.Fatal().Msgf("Error connecting to client: %s", err)
		}
		web3.client = c
	}

	// Bind Token
	tokenAddress := common.HexToAddress(address)

	if contractType == "ERC20" {
		t, err := ERC20.NewToken(tokenAddress, web3.client)
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
			log.Info().Msgf("🧩 Contract loaded: %s", address)
			// Add to DB
			if !preload {
				addContractToDB(newContract(address, contractType))
			}
			return true
		}
	}
	return false
}

func (web3 *Web3Client) reloadTokens() {
	// Load Tokens
	log.Debug().Msg("Loading Tokens From DB...")
	getContractsFromDB()
	for _, contract := range getContractsFromDB() {
		web3.loadToken(contract.Address, contract.Type, true)
	}
}
