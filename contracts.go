package main

import (
	"context"
	TOKEN "sync-ethers-go/abis/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

var token *TOKEN.Token

var ctx = context.Background()
var callOpts = &bind.CallOpts{Context: ctx, Pending: false}

func setupContracts() {
	log.Info().Msg("Setting up contracts...")
	loadABI()
	name, err := token.Name(callOpts)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get token name")
	}
	log.Info().Msgf("Loaded %s", name)
}

func loadABI() {
	log.Info().Msg("Loading ABI...")
	provider := "https://data-seed-prebsc-1-s1.binance.org:8545/"
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal().Msgf("Error connecting to client: %s", err)
	}
	tokenAddress := common.HexToAddress("4E0732efCdF0Cf92C48439535CD763de06FE353a")
	t, err := TOKEN.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal().Msgf("Some error occurred in TOKEN. Err: %s", err)
	}
	token = t
}
