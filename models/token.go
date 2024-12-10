package models

import (
	"errors"
	"math/big"
	"sync"
)

// Token represents the structure of a cryptocurrency token.
type Token struct {
	Symbol   string     // Symbol or name of the token
	Address  string     // Deployment address of the token
	Decimals int        // Number of decimals for the token
	Hold     bool       // Indicates if it's safe to hold this token
	GasFee   *big.Float // Latest gas fee in terms of Token.Symbol
}

// TokenList manages a collection of Tokens with efficient access methods.
type TokenList struct {
	symbolMap  map[string]*Token
	addressMap map[string]*Token
	mutex      sync.RWMutex
}

// NewTokenList initializes and returns a new TokenList.
func NewTokenList() *TokenList {
	return &TokenList{
		symbolMap:  make(map[string]*Token),
		addressMap: make(map[string]*Token),
	}
}

// NewTokenListFromSlice initializes and returns a new TokenList populated with tokens from a slice.
// Returns an error if there are duplicate symbols or addresses.
func NewTokenListFromSlice(tokens []*Token) (*TokenList, error) {
	tl := NewTokenList()
	tl.mutex.Lock()
	defer tl.mutex.Unlock()
	for _, token := range tokens {
		// Check for duplicate symbols
		if _, exists := tl.symbolMap[token.Symbol]; exists {
			return nil, errors.New("duplicate token symbol")
		}
		// Check for duplicate addresses
		if _, exists := tl.addressMap[token.Address]; exists {
			return nil, errors.New("duplicate token address")
		}
		// Add to maps
		tl.symbolMap[token.Symbol] = token
		tl.addressMap[token.Address] = token
	}
	return tl, nil
}

// AddToken adds a new token to the TokenList.
// It returns an error if a token with the same symbol or address already exists.
func (tl *TokenList) AddToken(token Token) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if _, exists := tl.symbolMap[token.Symbol]; exists {
		return errors.New("token with this symbol already exists")
	}
	if _, exists := tl.addressMap[token.Address]; exists {
		return errors.New("token with this address already exists")
	}

	// Create a copy to store pointers in maps
	t := token
	tl.symbolMap[token.Symbol] = &t
	tl.addressMap[token.Address] = &t

	return nil
}

// GetTokenBySymbol retrieves a token by its symbol.
// Returns an error if the token is not found.
func (tl *TokenList) GetTokenBySymbol(symbol string) (*Token, error) {
	tl.mutex.RLock()
	defer tl.mutex.RUnlock()

	token, exists := tl.symbolMap[symbol]
	if !exists {
		return nil, errors.New("no token found with the given symbol")
	}
	return token, nil
}

// GetTokenByAddress retrieves a token by its address.
// Returns an error if the token is not found.
func (tl *TokenList) GetTokenByAddress(address string) (*Token, error) {
	tl.mutex.RLock()
	defer tl.mutex.RUnlock()

	token, exists := tl.addressMap[address]
	if !exists {
		return nil, errors.New("no token found with the given address")
	}
	return token, nil
}

// UpdateTokenGasFeeBySymbol updates the gas fee in terms of a token identified by its symbol.
// Returns an error if the token is not found.
func (tl *TokenList) UpdateTokenGasFeeBySymbol(symbol string, gasFee *big.Float) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	token, exists := tl.symbolMap[symbol]
	if !exists {
		return errors.New("no token found with the given symbol")
	}
	token.GasFee = gasFee
	return nil
}

// RemoveTokenBySymbol removes a token from the TokenList by its symbol.
// Returns an error if the token is not found.
func (tl *TokenList) RemoveTokenBySymbol(symbol string) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	token, exists := tl.symbolMap[symbol]
	if !exists {
		return errors.New("no token found with the given symbol")
	}

	delete(tl.symbolMap, symbol)
	delete(tl.addressMap, token.Address)
	return nil
}

// ListTokens returns a slice of all tokens in the TokenList.
func (tl *TokenList) ListTokens() []*Token {
	tl.mutex.RLock()
	defer tl.mutex.RUnlock()

	tokens := make([]*Token, 0, len(tl.symbolMap))
	for _, token := range tl.symbolMap {
		tokens = append(tokens, token)
	}
	return tokens
}

// GetTokenSymbolList returns a list of all the symbols of tokens.
func (tl *TokenList) GetTokenSymbolList() []string {
	var symbolList []string

	allTokens := tl.ListTokens()
	for _, token := range allTokens {
		symbolList = append(symbolList, token.Symbol)
	}

	return symbolList
}
