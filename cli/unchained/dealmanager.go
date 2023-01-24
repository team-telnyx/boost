package unchained

import (
	"context"

	"github.com/filecoin-project/boost/storagemarket"
	lotusmarket "github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/lotus/api"
	ctypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func ChainDealManager() storagemarket.DealManagerIface {
	return &skipDM{}
}

type skipDM struct{}

func (s *skipDM) WaitForPublishDeals(ctx context.Context, publishCid cid.Cid, proposal market.DealProposal) (*lotusmarket.PublishDealsWaitResult, error) {
	return &lotusmarket.PublishDealsWaitResult{
		DealID:   0,
		FinalCid: publishCid,
	}, nil
}

func (s *skipDM) GetCurrentDealInfo(ctx context.Context, tok ctypes.TipSetKey, proposal *market.DealProposal, publishCid cid.Cid) (storagemarket.CurrentDealInfo, error) {
	return storagemarket.CurrentDealInfo{
		MarketDeal: &api.MarketDeal{
			Proposal: *proposal,
			State: market.DealState{
				SectorStartEpoch: abi.ChainEpoch(0),
			},
		},
	}, nil
}
