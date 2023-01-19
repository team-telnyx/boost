package unchained

import (
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)

type EmptyAsk struct {
}

func (e *EmptyAsk) GetAsk() *storagemarket.SignedStorageAsk {
	return &storagemarket.SignedStorageAsk{
		Ask: &storagemarket.StorageAsk{
			Price:         big.NewInt(0),
			VerifiedPrice: big.NewInt(0),
			MinPieceSize:  abi.PaddedPieceSize(0),
			MaxPieceSize:  abi.PaddedPieceSize(18446744073709551615),
		},
	}
}
