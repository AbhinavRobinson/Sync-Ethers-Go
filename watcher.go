package main

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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

// WatcherService service
type WatcherService struct {
	client *ethclient.Client
}

type WatcherConfig struct {
	ProviderURL string
}

// Watcher returns new service
func Watcher(input *WatcherConfig) *WatcherService {
	client, err := ethclient.Dial(input.ProviderURL)
	if err != nil {
		log.Fatal().Err(err)
	}

	return &WatcherService{
		client: client,
	}
}

// WatchOnChainTransactions watch for on-chain transactions
func (s *WatcherService) WatchOnChainTransactions(input []common.Address) error {
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
				header, err := s.client.HeaderByNumber(context.Background(), nil)
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

				logs, err := s.client.FilterLogs(context.Background(), query)
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

func RunWatcher() {
	watcher := Watcher(&WatcherConfig{
		ProviderURL: RPC,
	})

	err := watcher.WatchOnChainTransactions([]common.Address{
		common.HexToAddress("0x85d61e78d9062Cc7F9126CA9c2401bFcF7a4cF88"),
	})

	if err != nil {
		log.Fatal().Err(err)
	}
}
