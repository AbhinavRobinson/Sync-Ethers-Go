package main

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"
)

var web3 *Web3Client

func initApp() {
	log.Debug().Msg("Initializing App...")
	initMongo()
}

func initApi(host string, port int) {
	log.Debug().Msg("Initializing API...")
	setupApi(host, port)
}

func postInit() {
	log.Debug().Msg("Post-Initializing...")
	web3 = NewWeb3(&Web3Config{
		ProviderURL: RPC,
		CallOpts:    &bind.CallOpts{Context: context.Background(), Pending: false},
	})
	web3.reloadTokens()
}
