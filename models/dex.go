package models

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type DEXInstance interface {
	WatchPairSwaps(ethClient *ethclient.Client, pool *Pool, universalChan chan<- EventData)
}
