package strategy

import (
	"198/models"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Contains all opportunity information
type ArbitrageOpportunity struct {
	poolAB *models.Pool
	poolBC *models.Pool
	poolCA *models.Pool
	tokenA *models.Token
	tokenB *models.Token
	tokenC *models.Token
}

func ArbitrageStrategy(ethClient *ethclient.Client, tokenList *models.TokenList, poolList *models.PoolList) {
	tokenSymbols := tokenList.GetTokenSymbolList()
	minimumMultiplier := big.NewFloat(1) // Including gas fees

	// Search for arbitrage opportunities
	for _, tokenA := range tokenSymbols {
		tokenAInstance, err := tokenList.GetTokenBySymbol(tokenA)
		if err != nil {
			log.Printf("Failed to fetch TokenA: %v", err)
		}
		for _, tokenB := range tokenSymbols {
			if tokenA == tokenB {
				continue
			}
			for _, tokenC := range tokenSymbols {
				if tokenC == tokenA || tokenC == tokenB {
					continue
				}

				// Fetch pools
				poolAB, err := poolList.GetBestPoolForTokens(tokenA, tokenB)
				if err != nil {
					continue
				}
				poolBC, err := poolList.GetBestPoolForTokens(tokenB, tokenC)
				if err != nil {
					continue
				}
				poolCA, err := poolList.GetBestPoolForTokens(tokenC, tokenA)
				if err != nil {
					continue
				}

				// Fetch exchange rates
				rateAB, err := poolAB.AmountOutFromToken(tokenA) // tokenA -> tokenB
				if err != nil {
					continue
				}
				rateBC, err := poolBC.AmountOutFromToken(tokenB) // tokenB -> tokenC
				if err != nil {
					continue
				}
				rateCA, err := poolCA.AmountOutFromToken(tokenC) // tokenC -> tokenA
				if err != nil {
					continue
				}

				// Calculate cumulative exchange rate
				cumulativeExchangeRate := new(big.Float).Mul(rateAB, rateBC)
				cumulativeExchangeRate.Mul(cumulativeExchangeRate, rateCA)

				// Evaluate profitability
				if cumulativeExchangeRate.Cmp(minimumMultiplier) > 0 {
					// Profit opportunity detected
					log.Println("")
					log.Printf("Triangular Arbitrage opportunity: %v for %v -> %v -> %v -> %v", cumulativeExchangeRate, tokenA, tokenB, tokenC, tokenA)
					log.Printf("poolAB: %v", poolAB)
					log.Printf("poolBC: %v", poolBC)
					log.Printf("poolCA: %v", poolCA)

					// Construct opportunity object
					tokenBInstance, err := tokenList.GetTokenBySymbol(tokenB)
					if err != nil {
						log.Printf("Failed to fetch TokenB: %v", err)
					}
					tokenCInstance, err := tokenList.GetTokenBySymbol(tokenC)
					if err != nil {
						log.Printf("Failed to fetch TokenC: %v", err)
					}
					opp := ArbitrageOpportunity{
						poolAB: poolAB,
						poolBC: poolBC,
						poolCA: poolCA,
						tokenA: tokenAInstance,
						tokenB: tokenBInstance,
						tokenC: tokenCInstance,
					}

					// Save record of arbitrage opportunity
					log.Printf("Opportunity: %v", opp)
				}
			}
		}
	}
}
