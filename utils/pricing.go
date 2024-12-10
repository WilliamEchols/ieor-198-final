package utils

import (
	"math"
	"math/big"
)

// UniswapV3 CalculatePriceRatio computes the price ratio from sqrtPriceX96
func CalculatePriceRatio(sqrtPriceX96 *big.Int) (*big.Float, error) {
	sqrtPrice := new(big.Float).SetInt(sqrtPriceX96)
	factor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil))
	sqrtPriceQuotient := new(big.Float).Quo(sqrtPrice, factor)
	priceRatio := new(big.Float).Mul(sqrtPriceQuotient, sqrtPriceQuotient)

	return priceRatio, nil
}

// Adjust the price ratio considering the decimals of the tokens
func AdjustForTokenDecimals(priceRatio *big.Float, decimalsToken0, decimalsToken1 int) *big.Float {
	decimalsDifference := decimalsToken0 - decimalsToken1
	factor := new(big.Float).SetFloat64(math.Pow10(decimalsDifference))
	adjustedPrice := new(big.Float).Mul(priceRatio, factor)
	return adjustedPrice
}

func FeeToFeePercentage(fee *big.Int) float64 {
	feeFloat, _ := fee.Float64()
	feePercentage := feeFloat / 1e6
	return feePercentage
}
