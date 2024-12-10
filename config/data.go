package config

import (
	"198/models"
	"math/big"
)

var InitialTokens = []*models.Token{
	{ // 0
		Symbol:   "WBTC",
		Address:  "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6",
		Decimals: 6,
		Hold:     false,
		GasFee:   nil,
	},
	{ // 1
		Symbol:   "WETH",
		Address:  "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619",
		Decimals: 18,
		Hold:     false,
		GasFee:   nil,
	},
	{ // 2
		Symbol:   "USDC",
		Address:  "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359",
		Decimals: 6,
		Hold:     true,
		GasFee:   nil,
	},
	{ // 3
		Symbol:   "USDT",
		Address:  "0xc2132D05D31c914a87C6611C10748AEb04B58e8F",
		Decimals: 6,
		Hold:     true,
		GasFee:   nil,
	},
	{ // 4
		Symbol:   "WPOL",
		Address:  "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
		Decimals: 18,
		Hold:     false,
		GasFee:   nil,
	},
	{ // 5
		Symbol:   "DAI",
		Address:  "0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063",
		Decimals: 18,
		Hold:     true,
		GasFee:   nil,
	},
	{ // 6
		Symbol:   "LINK",
		Address:  "0x53E0bca35eC356BD5ddDFebbD1Fc0fD03FaBad39",
		Decimals: 18,
		Hold:     false,
		GasFee:   nil,
	},
}

// TODO- Add more pools
var InitialPools = []*models.Pool{
	{
		Address:               "0x50eaEDB835021E4A108B7290636d62E9765cc6d7",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[0], // WBTC
		Token1:                InitialTokens[1], // WETH
	},
	// TODO- WBTC scaling issue?
	{
		Address:               "0x32FAE204835e08b9374493d6B4628FD1F87DD045",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[0], // WBTC
		Token1:                InitialTokens[2], // USDC
	},
	{
		Address:               "0x167384319B41F7094e62f7506409Eb38079AbfF8",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(3000),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[1], // WETH
	},
	{
		Address:               "0xA4D8c89f0c20efbe54cBa9e7e7a7E509056228D9",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[2], // USDC
		Token1:                InitialTokens[1], // WETH
	},
	/*
		{
			Address:               "0xa374094527e1673a86de625aa59517c5de346d32",
			RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
			DEX:                   "UniswapV3",
			Fee:                   big.NewInt(500),
			Token0:                InitialTokens[4],   // WPOL
			Token1:                InitialTokens[100], // USDC.e
		},
	*/
	{
		Address:               "0x86f1d8390222a3691c28938ec7404a1661e618e0",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[1], // WETH
	},
	{
		Address:               "0x9b08288c3be4f62bbf8d1c20ac9c5e6f9467d8b7",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[3], // USDT
	},
	{
		Address:               "0xb6e57ed85c4c9dbfef2a68711e9d6f36c56e0fcb",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[2], // USDC
	},
	{
		Address:               "0x6b75F2189F0E11C52e814E09e280eb1a9A8A094a",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[0], // WBTC
	},
	{
		Address:               "0x0f663c16Dd7C65cF87eDB9229464cA77aEea536b",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[5], // DAI
	},
	{
		Address:               "0x0A28C2F5E0E8463E047C203F00F649812aE67E4f",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[6], // LINK
	},
	{
		Address:               "0x052C9b8f41f3855225495E78532aaAD0f22a925C",
		RouterContractAddress: "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
		DEX:                   "UniswapV3",
		Fee:                   big.NewInt(500),
		Token0:                InitialTokens[2], // USDC
		Token1:                InitialTokens[6], // LINK
	},
	{
		Address:               "0xdb975b96828352880409e86d5aE93c23c924f812",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(1081), // TODO- does fee change?
		Token0:                InitialTokens[0], // WBTC
		Token1:                InitialTokens[2], // USDC
	},
	{
		Address:               "0xa6AeDF7c4Ed6e821E67a6BfD56FD1702aD9a9719",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(888),
		Token0:                InitialTokens[2], // USDC
		Token1:                InitialTokens[1], // WETH
	},
	{
		Address:               "0x6669B4706cC152F359e947BCa68E263A87c52634",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(1419),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[2], // USDC
	},
	{
		Address:               "0xc10a06863f858f67C2Cd46F1675eE029D3F7acd8",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(14913),
		Token0:                InitialTokens[2], // USDC
		Token1:                InitialTokens[6], // LINK
	},
	{
		Address:               "0x479e1b71a702a595e19b6d5932cd5c863ab57ee0",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(900),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[1], // WETH
	},
	{
		Address:               "0xac4494e30a85369e332bdb5230d6d694d4259dbc",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(478),
		Token0:                InitialTokens[0], // WBTC
		Token1:                InitialTokens[1], // WETH
	},
	{
		Address:               "0x5b41eedcfc8e0ae47493d4945aa1ae4fe05430ff",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(1419),
		Token0:                InitialTokens[4], // WPOL
		Token1:                InitialTokens[3], // USDT
	},
	{
		Address:               "0x9ceff2f5138fc59eb925d270b8a7a9c02a1810f2",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(887),
		Token0:                InitialTokens[1], // WETH
		Token1:                InitialTokens[3], // USDT
	},
	{
		Address:               "0xab52931301078e2405c3a3ebb86e11ad0dfd2cfd",
		RouterContractAddress: "0xf5b509bB0909a69B1c207E495f687a596C168E12",
		DEX:                   "QuickswapV3",
		Fee:                   big.NewInt(2885),
		Token0:                InitialTokens[6], // LINK
		Token1:                InitialTokens[1], // WETH
	},
}
