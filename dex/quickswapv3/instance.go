package quickswapv3

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"198/models"
	"198/utils"
)

type Quickswapv3Instance struct {
	DEXSymbol string
}

func NewQuickswapV3Instance(symbol ...string) Quickswapv3Instance {
	instance := Quickswapv3Instance{
		DEXSymbol: "QuickswapV3", // default value
	}
	if len(symbol) > 0 && symbol[0] != "" {
		instance.DEXSymbol = symbol[0]
	}
	return instance
}

func (u Quickswapv3Instance) WatchPairSwaps(ethClient *ethclient.Client, pool *models.Pool, universalChan chan<- models.EventData) {
	DEXSymbol := u.DEXSymbol

	// local channel for direct subscription
	swapEventChan := make(chan *Quickswapv3Swap)

	// Specify the QuickswapV3 Pool contract address
	poolAddress := common.HexToAddress(pool.Address)

	// Create an instance of the pool contract
	poolContract, err := NewQuickswapv3(poolAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate %v contract: %v", DEXSymbol, err)
	}

	// Start watching for Swap events
	subscription, err := poolContract.WatchSwap(&bind.WatchOpts{
		Context: context.Background(),
	}, swapEventChan, nil, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to Swap events: %v", err)
	}
	defer subscription.Unsubscribe()
	log.Printf("[%v] Subscribed to Swap events (%v/%v) (pool: %v) (fee: %v)", DEXSymbol, pool.Token0.Symbol, pool.Token1.Symbol, poolAddress, pool.Fee)

	// Handle incoming Swap events
	for {
		select {
		case swapEvent := <-swapEventChan:
			// Calculate necessary info then redirect to a new universalChan

			// -- event latency --

			// Fetch the block header using the BlockNumber from the swap event
			blockHeader, err := ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(swapEvent.Raw.BlockNumber)))
			if err != nil {
				log.Printf("Failed to fetch block header: %v", err)
				continue
			}

			blockTimestamp := time.Unix(int64(blockHeader.Time), 0)
			currentTime := time.Now()
			latency := currentTime.Sub(blockTimestamp)

			// -- event latency --

			// Calculate the price ratio
			priceRatio, err := utils.CalculatePriceRatio(swapEvent.Price)
			if err != nil {
				log.Printf("Error calculating price ratio: %v", err)
				continue
			}

			// Output the exchange rate
			//adjustedPrice := utils.AdjustForTokenDecimals(priceRatio, pool.Token0.Decimals, pool.Token1.Decimals)
			//adjustedPrice := priceRatio
			adjustedPrice := utils.AdjustForTokenDecimals(priceRatio, pool.Token0.Decimals, pool.Token1.Decimals)

			// Calculate price for the other direction (Token0 per Token1)
			token0PerToken1 := new(big.Float).Quo(big.NewFloat(1), adjustedPrice)

			// Consider exchange fees and gas cost (assumed to be unitless)
			inputAmount := big.NewFloat(1.0) // Example input amount (1 token0)

			// Calculate fee amount
			feePercentage := utils.FeeToFeePercentage(pool.Fee)
			feeAmount := new(big.Float).Mul(inputAmount, big.NewFloat(feePercentage))

			// Net input amount after fee
			netInputAmount := new(big.Float).Sub(inputAmount, feeAmount)

			// Calculate output amount using net input amount and adjusted price
			amountOut := new(big.Float).Mul(netInputAmount, adjustedPrice)

			// Token1 -> Token0
			backwardsAmountsOut := new(big.Float).Mul(netInputAmount, token0PerToken1)

			// Parsed objet
			eventData := models.EventData{
				DEXSymbol:               DEXSymbol,
				PoolAddress:             pool.Address,
				BlockNumber:             swapEvent.Raw.BlockNumber,
				Latency:                 latency,
				Fee:                     pool.Fee,
				Token0Symbol:            pool.Token0.Symbol,
				Token1Symbol:            pool.Token1.Symbol,
				Token0ToToken1AmountOut: amountOut,
				Token1ToToken0AmountOut: backwardsAmountsOut,
			}

			// Send the structured data to the universal channel
			universalChan <- eventData
		case err := <-subscription.Err():
			log.Printf("ERROR: [%s] [%v] Subscription: %v", DEXSymbol, pool.Address, err)
		}
	}
}
