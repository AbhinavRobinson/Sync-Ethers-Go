package web3

import (
	ERC20 "sync-ethers-go/abis/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"

	config "sync-ethers-go/config"
	internal "sync-ethers-go/internal"
	mongo "sync-ethers-go/internal/mongo"
)

func (Web3 *Web3Client) LoadToken(address string, contractType string, preload bool) bool {
	log.Debug().Msg("Loading ABI...")

	// Dial Provider
	if Web3 == nil {
		provider := config.RPC
		c, err := ethclient.Dial(provider)
		if err != nil {
			log.Fatal().Msgf("Error connecting to client: %s", err)
		}
		Web3.client = c
	}

	// Bind Token
	tokenAddress := common.HexToAddress(address)

	if contractType == "ERC20" {
		t, err := ERC20.NewToken(tokenAddress, Web3.client)
		if err != nil {
			log.Fatal().Msgf("Some error occurred in TOKEN. Err: %s", err)
		}
		log.Info().Msg("🧩 Contract processing...")
		// Add to mapping
		if internal.Contracts.Tokens == nil {
			// Allocate if not found
			internal.Contracts.Tokens = make(map[string]*ERC20.Token)
		}
		if internal.Contracts.Tokens[address] == nil {
			// Add if not found
			internal.Contracts.Tokens[address] = t
			log.Info().Msgf("🧩 Contract loaded: %s", address)
			// Add to DB
			if !preload {
				mongo.AddContractToDB(mongo.NewContract(address, contractType))
			}
			return true
		}
	}
	return false
}

func (Web3 *Web3Client) ReloadTokens() {
	// Load Tokens
	log.Debug().Msg("Loading Tokens From DB...")
	mongo.GetContractsFromDB()
	for _, contract := range mongo.GetContractsFromDB() {
		Web3.LoadToken(contract.Address, contract.Type, true)
	}
}
