package models

import (
	"errors"
	"math/big"
	"sync"
)

// NOTE: pointing to Token so that we can modify token balances and see that reflected from a PoolList search
type Pool struct {
	Address                 string
	DEX                     string
	RouterContractAddress   string
	Fee                     *big.Int
	Token0                  *Token
	Token1                  *Token
	Token0ToToken1AmountOut *big.Float
	Token1ToToken0AmountOut *big.Float
}

// PoolList manages a collection of Pools with efficient access methods.
type PoolList struct {
	symbolMap  map[string]*Pool
	addressMap map[string]*Pool
	mutex      sync.RWMutex
}

// NewPoolList initializes and returns a new PoolList.
func NewPoolList() *PoolList {
	return &PoolList{
		symbolMap:  make(map[string]*Pool),
		addressMap: make(map[string]*Pool),
	}
}

// NewPoolListFromSlice initializes and returns a new PoolList populated with pools from a slice.
// Returns an error if there are duplicate symbols or addresses.
func NewPoolListFromSlice(pools []*Pool) (*PoolList, error) {
	pl := NewPoolList()
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	for _, pool := range pools {
		if _, exists := pl.addressMap[pool.Address]; exists {
			return nil, errors.New("duplicate pool address")
		}

		pl.addressMap[pool.Address] = pool
	}
	return pl, nil
}

// AddPool adds a new pool to the PoolList.
// It returns an error if a pool with the same address already exists.
func (pl *PoolList) AddPool(pool Pool) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	if _, exists := pl.addressMap[pool.Address]; exists {
		return errors.New("pool with this address already exists")
	}

	// Create a copy to store pointers in maps
	p := pool
	pl.addressMap[pool.Address] = &p

	return nil
}

// GetPoolByAddress retrieves a pool by its address.
// Returns an error if the pool is not found.
func (pl *PoolList) GetPoolByAddress(address string) (*Pool, error) {
	pl.mutex.RLock()
	defer pl.mutex.RUnlock()

	pool, exists := pl.addressMap[address]
	if !exists {
		return nil, errors.New("no pool found with the given address")
	}
	return pool, nil
}

// GetPoolByTokenAndDexSymbols retrieves a pool by the symbols of its token and the symbol of its dex.
// Returns an error if the pool is not found.
func (pl *PoolList) GetPoolByTokenAndDexSymbols(symbol0, symbol1, dex string) (*Pool, error) {
	pl.mutex.RLock()
	defer pl.mutex.RUnlock()

	allPools := pl.ListPools()
	for _, pool := range allPools {
		if pool.DEX != dex {
			continue
		}

		if (pool.Token0.Symbol == symbol0 && pool.Token1.Symbol == symbol1) || (pool.Token0.Symbol == symbol1 && pool.Token1.Symbol == symbol0) {
			return pool, nil
		}
	}

	return nil, errors.New("not pool with given token symbols and dex symbol could be found")
}

// RemovePoolByAddress removes a pool from the PoolList by its address.
// Returns an error if the token is not found.
func (pl *PoolList) RemovePoolByAddress(address string) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	pool, exists := pl.addressMap[address]
	if !exists {
		return errors.New("no pool found with the given symbol")
	}

	delete(pl.addressMap, pool.Address)
	return nil
}

// ListPools returns a slice of all pools in the PoolList.
func (pl *PoolList) ListPools() []*Pool {
	pl.mutex.RLock()
	defer pl.mutex.RUnlock()

	pools := make([]*Pool, 0, len(pl.addressMap))
	for _, pool := range pl.addressMap {
		pools = append(pools, pool)
	}
	return pools
}

// UpdatePoolAmountOutsByAddress takes in an address (to specify a pool) then token0ToToken1AmountOut, token1ToToken0AmountOut to update for that pool.
func (pl *PoolList) UpdatePoolAmountOutsByAddress(address string, token0ToToken1AmountOut, token1ToToken0AmountOut *big.Float) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	pool, exists := pl.addressMap[address]
	if !exists {
		return errors.New("no pool found with the given address")
	}
	pool.Token0ToToken1AmountOut = token0ToToken1AmountOut
	pool.Token1ToToken0AmountOut = token1ToToken0AmountOut
	return nil
}

// GetBestPoolForTokens finds the best pool to convert from one token to another.
// It returns the pool that provides the maximum amount out for the given token pair.
func (pl *PoolList) GetBestPoolForTokens(fromTokenSymbol, toTokenSymbol string) (*Pool, error) {
	pl.mutex.RLock()
	defer pl.mutex.RUnlock()

	var bestPool *Pool
	maxAmountOut := big.NewFloat(0)

	for _, pool := range pl.addressMap {
		if pool.Token0ToToken1AmountOut == nil || pool.Token1ToToken0AmountOut == nil {
			continue
		}
		if pool.Token0.Symbol == fromTokenSymbol && pool.Token1.Symbol == toTokenSymbol {
			// Direction is Token0 to Token1
			amountOut := pool.Token0ToToken1AmountOut
			if amountOut.Cmp(maxAmountOut) > 0 {
				maxAmountOut = amountOut
				bestPool = pool
			}
		} else if pool.Token1.Symbol == fromTokenSymbol && pool.Token0.Symbol == toTokenSymbol {
			// Direction is Token1 to Token0
			amountOut := pool.Token1ToToken0AmountOut
			if amountOut.Cmp(maxAmountOut) > 0 {
				maxAmountOut = amountOut
				bestPool = pool
			}
		}
	}

	if bestPool == nil {
		return nil, errors.New("no pool found for given tokens")
	}

	return bestPool, nil
}

// AmountOutFromToken returns the amount out when swapping from the given token.
// If the token is not part of the pool, it returns an error.
func (p *Pool) AmountOutFromToken(tokenSymbol string) (*big.Float, error) {
	if p.Token0.Symbol == tokenSymbol {
		return p.Token0ToToken1AmountOut, nil
	} else if p.Token1.Symbol == tokenSymbol {
		return p.Token1ToToken0AmountOut, nil
	} else {
		return nil, errors.New("token not found in pool")
	}
}

// AmountOutToToken returns the amount out when swapping to the given token.
// If the token is not part of the pool, it returns an error.
func (p *Pool) AmountOutToToken(tokenSymbol string) (*big.Float, error) {
	if p.Token0.Symbol == tokenSymbol {
		// Swapping to Token0: amount out is Token1ToToken0AmountOut
		return p.Token1ToToken0AmountOut, nil
	} else if p.Token1.Symbol == tokenSymbol {
		// Swapping to Token1: amount out is Token0ToToken1AmountOut
		return p.Token0ToToken1AmountOut, nil
	} else {
		return nil, errors.New("token not found in pool")
	}
}
