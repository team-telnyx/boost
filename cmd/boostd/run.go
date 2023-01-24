package main

import (
	"fmt"

	"github.com/filecoin-project/boost/api"
	"github.com/filecoin-project/boost/cli/unchained"
	"github.com/filecoin-project/boost/indexprovider"
	"github.com/filecoin-project/boost/node"
	"github.com/filecoin-project/boost/node/modules"
	"github.com/filecoin-project/boost/node/modules/dtypes"
	boostmarket "github.com/filecoin-project/boost/storagemarket"
	"github.com/filecoin-project/boost/storagemarket/types"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/libp2p/go-libp2p/core/peer"

	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/api/v1api"
	lcli "github.com/filecoin-project/lotus/cli"
	lotus_storageadapter "github.com/filecoin-project/lotus/markets/storageadapter"
	lotus_modules "github.com/filecoin-project/lotus/node/modules"
	lotus_dtypes "github.com/filecoin-project/lotus/node/modules/dtypes"
	lotus_repo "github.com/filecoin-project/lotus/node/repo"
	"github.com/filecoin-project/lotus/storage/sealer"

	"net/http"
	_ "net/http/pprof"

	"github.com/urfave/cli/v2"
)

var runCmd = &cli.Command{
	Name:   "run",
	Usage:  "Start a boost process",
	Before: before,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "pprof",
			Usage: "run pprof web server on localhost:6060",
		},
		&cli.BoolFlag{
			Name:  "nosync",
			Usage: "dont wait for the full node to sync with the chain",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Bool("pprof") {
			go func() {
				err := http.ListenAndServe("localhost:6060", nil)
				if err != nil {
					log.Error(err)
				}
			}()
		}

		fullnodeApi, ncloser, err := lcli.GetFullNodeAPIV1(cctx)
		if err != nil {
			return fmt.Errorf("getting full node api: %w", err)
		}
		defer ncloser()

		ctx := lcli.ReqContext(cctx)

		log.Debug("Checking full node version")

		v, err := fullnodeApi.Version(ctx)
		if err != nil {
			return err
		}

		log.Debugw("Remote full node version", "version", v)

		if !v.APIVersion.EqMajorMinor(lapi.FullAPIVersion1) {
			return fmt.Errorf("Remote API version didn't match (expected %s, remote %s)", lapi.FullAPIVersion1, v.APIVersion)
		}

		log.Debug("Checking full node sync status")

		if !cctx.Bool("nosync") {
			if err := lcli.SyncWait(ctx, &v0api.WrapperV1Full{FullNode: fullnodeApi}, false); err != nil {
				return fmt.Errorf("sync wait: %w", err)
			}
		}

		boostRepoPath := cctx.String(FlagBoostRepo)

		r, err := lotus_repo.NewFS(boostRepoPath)
		if err != nil {
			return err
		}
		ok, err := r.Exists()
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("repo at '%s' is not initialized", cctx.String(FlagBoostRepo))
		}

		shutdownChan := make(chan struct{})

		log.Debug("Instantiating new boost node")

		var boostApi api.Boost
		options := []node.Option{
			node.BoostAPI(&boostApi),
			node.Override(new(dtypes.ShutdownChan), shutdownChan),
			node.Base(),
			node.Repo(r),
			node.Override(new(v1api.FullNode), fullnodeApi),
		}
		if cctx.Bool(FlagUnchained) {
			log.Infow("Not supporting paid deals when running without a chain")

			options = append(options, []node.Option{
				node.Override(node.HandleDealsKey, func() error { return nil }),
				node.Override(new(sealer.StorageAuth), lotus_modules.StorageAuth),
				node.Override(new(*indexprovider.Wrapper), indexprovider.NewWrapperNoLegacy()),
				node.Override(node.HandleBoostDealsKey, modules.HandleOnlyBoostDeals),
				node.Override(new(types.AskGetter), func() types.AskGetter { return &unchained.EmptyAsk{} }),
				node.Override(new(storagemarket.StorageProvider), func() storagemarket.StorageProvider { return nil }),
				node.Override(new(*lotus_storageadapter.DealPublisher), unchained.DealPublishAccepter(cctx)),
				node.Override(new(boostmarket.DealManagerIface), unchained.ChainDealManager),
				node.Override(new(lotus_dtypes.MinerAddress), unchained.MinerAddress),
			}...)
		}
		stop, err := node.New(ctx, options...)
		if err != nil {
			return fmt.Errorf("creating node: %w", err)
		}

		log.Debug("Getting API endpoint of boost node")

		endpoint, err := r.APIEndpoint()
		if err != nil {
			return fmt.Errorf("getting API endpoint: %w", err)
		}

		// Get maddr for boost node
		maddr, err := boostApi.NetAddrsListen(ctx)
		if err != nil {
			return fmt.Errorf("getting boost libp2p address: %w", err)
		}

		log.Infow("Boost libp2p node listening", "maddr", maddr)

		// Bootstrap with full node
		remoteAddrs, err := fullnodeApi.NetAddrsListen(ctx)
		if err != nil {
			return fmt.Errorf("getting full node libp2p address: %w", err)
		}

		log.Debugw("Bootstrapping boost libp2p network with full node", "maadr", remoteAddrs)

		if remoteAddrs.ID != peer.ID("") {
			if err := boostApi.NetConnect(ctx, remoteAddrs); err != nil {
				return fmt.Errorf("connecting to full node (libp2p): %w", err)
			}
		}

		// Instantiate the boost service JSON RPC handler.
		handler, err := node.BoostHandler(boostApi, true)
		if err != nil {
			return fmt.Errorf("failed to instantiate rpc handler: %w", err)
		}

		log.Infow("Boost JSON RPC server is listening", "endpoint", endpoint)

		// Serve the RPC.
		rpcStopper, err := node.ServeRPC(handler, "boost", endpoint)
		if err != nil {
			return fmt.Errorf("failed to start json-rpc endpoint: %s", err)
		}

		// Monitor for shutdown.
		finishCh := node.MonitorShutdown(shutdownChan,
			node.ShutdownHandler{Component: "rpc server", StopFunc: rpcStopper},
			node.ShutdownHandler{Component: "boost", StopFunc: stop},
		)

		<-finishCh
		return nil
	},
}
