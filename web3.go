package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

// Web3Client service
type Web3Client struct {
	client   *ethclient.Client
	CallOpts *bind.CallOpts
}

type Web3Config struct {
	ProviderURL string
	CallOpts    *bind.CallOpts
}

// NewWeb3 returns new service
func NewWeb3(input *Web3Config) *Web3Client {
	client, err := ethclient.Dial(input.ProviderURL)
	if err != nil {
		log.Fatal().Err(err)
	}

	return &Web3Client{
		client:   client,
		CallOpts: input.CallOpts,
	}
}
