package entity

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type SubmittedBlocks struct {
	
}

type SubmittedBlock struct {
	BlockType uint8 `json:"blockType"`
	BlockSize uint16 `json:"blockSize"`
	BlockVersion uint8 `json:"blockVersion"`
	Data []byte `json:"data"`
	Proof [8]*big.Int `json:"proof"`
	StoreBlockInfoOnchain bool `json:"storeBlockInfoOnchain"`
	AuxiliaryData []byte `json:"auxiliaryData"`
	OffchainData []byte `json:"offchainData"`
}

// type AuxiliaryData struct {
// 	Data []*Auxiliary
// }
type AuxiliaryData struct {
	Data []Auxiliary
 }

type Auxiliary struct {
	Data []byte
 }

type WithdrawAuxData struct {
	StoreRecipient bool
	GasLimit       *big.Int
	Signature      []byte
	MinGas         *big.Int
	To             common.Address
	MaxFee         *big.Int
	ValidUntil     uint32
	Amount         *big.Int
}
