package dex

import (
	"198/dex/quickswapv3"
	"198/dex/uniswapv3"
	"198/models"
)

// DEXImplementations holds all the DEX instances.
var DEXImplementations = map[string]models.DEXInstance{
	"UniswapV3":   uniswapv3.NewUniswapV3Instance(),
	"SushiswapV3": uniswapv3.NewUniswapV3Instance("SushiswapV3"), // NOTE: Ensure we are using the correct router in the implementation
	"QuickswapV3": quickswapv3.NewQuickswapV3Instance(),
}
