package main

import (
	"context"
	TOKEN "sync-ethers-go/abis/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

var contracts Contracts

var ctx = context.Background()
var callOpts = &bind.CallOpts{Context: ctx, Pending: false}
var client *ethclient.Client

func loadToken(address string) bool {
	log.Debug().Msg("Loading ABI...")

	// Dial Provider
	if client == nil {
		provider := "https://data-seed-prebsc-1-s1.binance.org:8545/"
		c, err := ethclient.Dial(provider)
		if err != nil {
			log.Fatal().Msgf("Error connecting to client: %s", err)
		}
		client = c
	}

	// Bind Token
	tokenAddress := common.HexToAddress(address)
	t, err := TOKEN.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal().Msgf("Some error occurred in TOKEN. Err: %s", err)
	}
	log.Info().Msg("🧩 Contract processed.")

	// Add to mapping
	if contracts.Tokens == nil {
		contracts.Tokens = make(map[string]*TOKEN.Token)
	}
	contracts.Tokens[address] = t
	return true
}
