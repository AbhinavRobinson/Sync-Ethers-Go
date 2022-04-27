package main

import TOKEN "sync-ethers-go/abis/token"

type Contracts struct {
	Tokens map[string]*TOKEN.Token
}

var AvailableContractTypes = map[string]bool{
	"ERC20":   true,
	"ERC721":  false,
	"ERC1155": false,
}
