package unchained

import (
	"context"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/v1api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/libp2p/go-libp2p-core/peer"
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
	return peer.AddrInfo{}, nil
}
