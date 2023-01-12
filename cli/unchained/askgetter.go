package unchained

import "github.com/filecoin-project/go-fil-markets/storagemarket"

type EmptyAsk struct {
}

func (e *EmptyAsk) GetAsk() *storagemarket.SignedStorageAsk {
	return nil
}
