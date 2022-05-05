package api

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"

	config "sync-ethers-go/config"
	mongo "sync-ethers-go/internal/mongo"
	W3 "sync-ethers-go/internal/web3"
)

func InitApp() {
	log.Debug().Msg("Initializing App...")
	mongo.InitMongo()
}

func InitApi(host string, port int) {
	log.Debug().Msg("Initializing API...")
	SetupApi(host, port)
}

func PostInit() {
	log.Debug().Msg("Post-Initializing...")
	W3.Web3 = W3.NewWeb3(&W3.Web3Config{
		ProviderURL: config.RPC,
		CallOpts:    &bind.CallOpts{Context: context.Background(), Pending: false},
	})
	W3.Web3.ReloadTokens()
}
