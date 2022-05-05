package internal

import (
	ERC20 "sync-ethers-go/abis/erc20"

	"github.com/kamva/mgm/v3"
)

var Contracts ContractList

type ContractList struct {
	Tokens map[string]*ERC20.Token
}

var AvailableContractTypes = map[string]bool{
	"ERC20":   true,
	"ERC721":  false,
	"ERC1155": false,
}

type StoredContracts struct {
	mgm.DefaultModel `bson:",inline"`
	Address          string `json:"address"`
	Type             string `json:"type"`
}
