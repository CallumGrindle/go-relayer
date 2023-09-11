package services

import (
	"api/internal/config"
	"log"

	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var blockchainClient *ethclient.Client
var onceBlockchainClient sync.Once

func getBlockchainClient() *ethclient.Client {
	var err error

	onceBlockchainClient.Do(func() {
		log.Printf("Initializing blockchain connection: %v", config.ApplicationConfig.RPC_URL)
		blockchainClient, err = ethclient.Dial(config.ApplicationConfig.RPC_URL)

		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
	})

	return blockchainClient
}
