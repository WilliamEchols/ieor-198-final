package models

import (
	"math/big"
	"time"
)

type EventData struct {
	DEXSymbol               string
	PoolAddress             string
	BlockNumber             uint64
	Latency                 time.Duration
	Fee                     *big.Int
	Token0Symbol            string
	Token1Symbol            string
	Token0ToToken1AmountOut *big.Float
	Token1ToToken0AmountOut *big.Float
}
