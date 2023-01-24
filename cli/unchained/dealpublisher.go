package unchained

import (
	"time"

	lotus_storageadapter "github.com/filecoin-project/lotus/markets/storageadapter"
	"github.com/filecoin-project/lotus/storage/ctladdr"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

func DealPublishAccepter(cctx *cli.Context) func(lc fx.Lifecycle, as *ctladdr.AddressSelector) *lotus_storageadapter.DealPublisher {
	return func(lc fx.Lifecycle, as *ctladdr.AddressSelector) *lotus_storageadapter.DealPublisher {
		f := lotus_storageadapter.NewDealPublisher(nil, lotus_storageadapter.PublishMsgConfig{
			Period:                  time.Hour,
			MaxDealsPerMsg:          1, // todo: from cfg.LotusDealmaking.MaxDealsPerPublishMsg,
			StartEpochSealingBuffer: 0, // todo: from cfg.LotusDealmaking.StartEpochSealingBuffer,
		})
		// the accepter calls the API to publish. we just accept that instead.
		fn := MakeProxy(cctx)
		return f(lc, fn, as)
	}
}
