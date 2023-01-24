package unchained

import (
	"context"
	"fmt"

	"github.com/filecoin-project/boost/sealingpipeline"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/filecoin-project/lotus/chain/types"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	lotus_dtypes "github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/urfave/cli/v2"
)

type Node struct {
	v1api.FullNode

	closer jsonrpc.ClientCloser
}

func (n *Node) Close() {
	n.closer()
}

func MakeProxy(cctx *cli.Context) v1api.FullNode {
	proxy, closer, err := cliutil.GetFullNodeAPIV1(cctx)
	if err != nil {
		panic(err)
	}
	return &Node{
		FullNode: proxy,

		closer: closer,
	}
}

// Over-ridden API methods:
func (n *Node) NetAddrsListen(ctx context.Context) (peer.AddrInfo, error) {
	return peer.AddrInfo{
		ID: peer.ID(""),
	}, nil
}

func (n *Node) NetProtectAdd(ctx context.Context, pids []peer.ID) error {
	return fmt.Errorf("not supported")
}

func (n *Node) StateMinerInfo(ctx context.Context, actor address.Address, tsk types.TipSetKey) (api.MinerInfo, error) {
	return api.MinerInfo{
		Owner:            actor,
		Worker:           actor,
		NewWorker:        actor,
		ControlAddresses: []address.Address{actor},
	}, nil
}

func MinerAddress(ds lotus_dtypes.MetadataDS, spapi sealingpipeline.API) (lotus_dtypes.MinerAddress, error) {
	// first try to get from the API...
	apiAddr, err := spapi.ActorAddress(context.TODO())
	if err == nil {
		return lotus_dtypes.MinerAddress(apiAddr), nil
	}

	maddrb, err := ds.Get(context.TODO(), datastore.NewKey("miner-address"))
	if err != nil {
		return lotus_dtypes.MinerAddress(address.Undef), err
	}
	addr, _ := address.NewFromBytes(maddrb)
	return lotus_dtypes.MinerAddress(addr), err
}
