package web3

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	ERC20 "sync-ethers-go/abis/erc20"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

// WatchOnChainTransactions watch for on-chain transactions
func (Web3 *Web3Client) WatchOnChainTransactions(input []common.Address) error {
	contractAbi, err := abi.JSON(strings.NewReader(string(ERC20.TokenABI)))
	if err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		currentBlockNumber := big.NewInt(0)
		lastBlockNumberScanned := big.NewInt(0)
		for {
			select {
			case <-ticker.C:
				header, err := Web3.client.HeaderByNumber(context.Background(), nil)
				if err != nil {
					log.Error().Msgf("error HeaderByNumber: %s", err)
					continue
				}

				// if block number is the same then just continue
				if header.Number.Cmp(currentBlockNumber) == 0 {
					continue
				}

				// start a little behind to make sure we didn't miss any txs
				currentBlockNumber = new(big.Int).Sub(header.Number, big.NewInt(2))
				if lastBlockNumberScanned.Cmp(currentBlockNumber) == 0 {
					continue
				}

				log.Debug().Msgf("scanning block: %v", currentBlockNumber)
				lastBlockNumberScanned = currentBlockNumber
				query := ethereum.FilterQuery{
					FromBlock: currentBlockNumber,
					ToBlock:   nil,
					Addresses: input,
				}

				logs, err := Web3.client.FilterLogs(context.Background(), query)
				if err != nil {
					log.Error().Msgf("error getting FilterLogs: %s", err)
					continue
				}

				for _, vLog := range logs {
					event, err := contractAbi.Unpack("Transfer", vLog.Data)
					if err != nil {
						log.Fatal().Err(err)
					}
					fmt.Println(event)
					log.Debug().Msgf("Logger: %v", vLog)
				}
			}
		}
	}()

	return nil
}

func StartWatcher(addresses []common.Address) {
	log.Info().Msg("Starting Watcher...")
	for _, addresses := range addresses {
		log.Info().Msgf("Watching: %s", addresses)
	}

	err := Web3.WatchOnChainTransactions(addresses)

	if err != nil {
		log.Fatal().Err(err)
	}
}
