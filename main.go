package main

import (
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"198/config"
	"198/dex"
	"198/models"
	"198/strategy"
	"198/utils"
)

func main() {
	// Setup logging
	logFile, err := utils.SetupLogging()
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Load environment variables
	err = godotenv.Load(".env.polygon")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	NODE_URL := os.Getenv("NODE_URL")
	NODE_NAME := os.Getenv("NODE_NAME")

	// Connect to the Polygon node (WSS)
	ethClient, err := ethclient.Dial(NODE_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the node via ethclient: %v", err)
	}
	defer ethClient.Close()
	log.Printf("Connected to %v node", NODE_NAME)

	// Instantiate tokenList
	initialTokens := config.InitialTokens
	tokenList, err := models.NewTokenListFromSlice(initialTokens)
	if err != nil {
		log.Fatalf("Failed to construct tokenList: %v", err)
	}

	// Centralized (universal) channel for aggregate events
	universalChan := make(chan models.EventData)

	// Instantiate pools
	initialPools := config.InitialPools
	poolList, err := models.NewPoolListFromSlice(initialPools)
	if err != nil {
		log.Printf("Failed to construct tokenList: %v", err)
	}

	// Listen to pools on each DEX (w/ delay to avoid high API/s use for initialization)
	allPools := poolList.ListPools()
	for _, pool := range allPools {
		dexImpl, ok := dex.DEXImplementations[pool.DEX]
		if !ok {
			log.Fatalf("DEX implementation for %s not found", pool.DEX)
		}

		// Start listener for the pool
		go dexImpl.WatchPairSwaps(ethClient, pool, universalChan)

		// Wait for 0.5 seconds before starting the next listener (prevent spamming network)
		time.Sleep(500 * time.Millisecond)
	}

	// Process parsed events
	for event := range universalChan {
		// Update rates for associated pool
		poolList.UpdatePoolAmountOutsByAddress(event.PoolAddress, event.Token0ToToken1AmountOut, event.Token1ToToken0AmountOut)

		// Log swap event info
		log.Printf("New swap event: %v", event)

		// Identify and log triangular arbitrage opportunities
		strategy.ArbitrageStrategy(ethClient, tokenList, poolList)
	}
}
