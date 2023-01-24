package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/filecoin-project/go-address"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-commp-utils/writer"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	ctypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet/key"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/lib/sigs"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car/v2"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/urfave/cli/v2"

	bcli "github.com/filecoin-project/boost/cli"
	clinode "github.com/filecoin-project/boost/cli/node"
	"github.com/filecoin-project/boost/cmd"
	"github.com/filecoin-project/boost/storagemarket/types"
)

const DealProtocolv120 = "/fil/storage/mk/1.2.0"

var dealCmd = &cli.Command{
	Name:        "deal",
	Description: "Make a direct deal with a running boost market node",
	Before:      before,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "market-provider",
			Aliases:  []string{"mp"},
			Usage:    "multiaddr of the node to make a dela with",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := bcli.ReqContext(cctx)

		// for getting current chain epoch
		api, closer, err := lcli.GetGatewayAPI(cctx)
		if err != nil {
			return fmt.Errorf("cant setup gateway connection: %w", err)
		}
		defer closer()
		var startEpoch abi.ChainEpoch
		if cctx.IsSet("start-epoch") {
			startEpoch = abi.ChainEpoch(cctx.Int("start-epoch"))
		} else {
			tipset, err := api.ChainHead(ctx)
			if err != nil {
				return fmt.Errorf("getting chain head: %w", err)
			}

			head := tipset.Height()

			log.Debugw("current block height", "number", head)

			startEpoch = head + abi.ChainEpoch(5760) // head + 2 days
		}

		n, err := clinode.Setup(cctx.String(cmd.FlagRepo.Name))
		if err != nil {
			return err
		}

		ma, err := multiaddr.NewMultiaddr(cctx.String("market-provider"))
		if err != nil {
			return err
		}
		ai, err := peer.AddrInfoFromP2pAddr(ma)
		if err != nil {
			return err
		}
		if err := n.Host.Connect(ctx, *ai); err != nil {
			return fmt.Errorf("failed to connect to peer %s: %w", ai.ID, err)
		}

		x, err := n.Host.Peerstore().FirstSupportedProtocol(ai.ID, DealProtocolv120)
		if err != nil {
			return fmt.Errorf("getting protocols: %w", err)
		}

		if len(x) == 0 {
			return fmt.Errorf("cannot make a deal with markets endpoint")
		}

		carFile := cctx.Args().First()
		cfStat, err := os.Stat(carFile)
		if err != nil {
			return err
		}
		carReader, err := car.OpenReader(carFile)
		if err != nil {
			return err
		}
		roots, err := carReader.Roots()
		if err != nil {
			return err
		}

		rdr, err := os.Open(carFile)
		if err != nil {
			return err
		}
		defer rdr.Close() //nolint:errcheck

		w := &writer.Writer{}
		_, err = io.CopyBuffer(w, rdr, make([]byte, writer.CommPBuf))
		if err != nil {
			return fmt.Errorf("copy into commp writer: %w", err)
		}

		commP, err := w.Sum()
		if err != nil {
			return fmt.Errorf("computing commP failed: %w", err)
		}

		randomMinerKey, _ := key.GenerateKey(ctypes.KTSecp256k1)
		err = TryProposalWith(cctx, n.Host, ai.ID, roots[0], commP.PieceCID, startEpoch, cfStat.Size(), randomMinerKey.Address)
		if err != nil {
			rexp := regexp.MustCompile(`.*incorrect provider for deal.*provider\.Address: ([\w]+)`)
			if m := rexp.FindStringSubmatch(err.Error()); len(m) > 1 {
				realMiner, _ := address.NewFromString(m[1])
				err = TryProposalWith(cctx, n.Host, ai.ID, roots[0], commP.PieceCID, startEpoch, cfStat.Size(), realMiner)
			}
		}
		return err
	},
}

func TryProposalWith(cctx *cli.Context,
	h host.Host,
	peer peer.ID,
	root cid.Cid,
	commP cid.Cid,
	startEpoch abi.ChainEpoch,
	size int64,
	providerAddr address.Address) error {
	ctx := cctx.Context
	localKey, _ := key.GenerateKey(ctypes.KTSecp256k1)
	dealUuid := uuid.New()

	transfer := types.Transfer{
		Size: uint64(size),
		Type: "manual",
	}

	paddedSize := abi.PaddedPieceSize(NextPowOf2(int(size)))

	dealProposal, err := dealProposal(ctx,
		localKey,
		localKey.Address,
		root,
		paddedSize,
		commP,
		providerAddr,
		startEpoch,
		518400, // duration
		false,
		abi.NewTokenAmount(0), //providerCollateral
		abi.NewTokenAmount(0), //storagePrice
	)
	if err != nil {
		return fmt.Errorf("failed to create a deal proposal: %w", err)
	}

	dealParams := types.DealParams{
		DealUUID:           dealUuid,
		ClientDealProposal: *dealProposal,
		DealDataRoot:       root,
		IsOffline:          true,
		Transfer:           transfer,
		RemoveUnsealedCopy: cctx.Bool("remove-unsealed-copy"),
		SkipIPNIAnnounce:   cctx.Bool("skip-ipni-announce"),
	}

	s, err := h.NewStream(ctx, peer, DealProtocolv120)
	if err != nil {
		return fmt.Errorf("failed to open stream: %w", err)
	}
	defer s.Close()

	var resp types.DealResponse
	if err := doRpc(ctx, s, &dealParams, &resp); err != nil {
		return fmt.Errorf("send proposal rpc: %w", err)
	}

	if !resp.Accepted {
		return fmt.Errorf("deal proposal rejected: %s", resp.Message)
	}

	return nil
}

func dealProposal(
	ctx context.Context,
	k *key.Key,
	clientAddr address.Address,
	rootCid cid.Cid,
	pieceSize abi.PaddedPieceSize,
	pieceCid cid.Cid,
	minerAddr address.Address,
	startEpoch abi.ChainEpoch,
	duration int,
	verified bool,
	providerCollateral abi.TokenAmount,
	storagePrice abi.TokenAmount,
) (*market.ClientDealProposal, error) {
	endEpoch := startEpoch + abi.ChainEpoch(duration)
	// deal proposal expects total storage price for deal per epoch, therefore we
	// multiply pieceSize * storagePrice (which is set per epoch per GiB) and divide by 2^30
	storagePricePerEpochForDeal := big.Div(big.Mul(big.NewInt(int64(pieceSize)), storagePrice), big.NewInt(int64(1<<30)))
	l, err := market.NewLabelFromString(rootCid.String())
	if err != nil {
		return nil, err
	}
	proposal := market.DealProposal{
		PieceCID:             pieceCid,
		PieceSize:            pieceSize,
		VerifiedDeal:         verified,
		Client:               clientAddr,
		Provider:             minerAddr,
		Label:                l,
		StartEpoch:           startEpoch,
		EndEpoch:             endEpoch,
		StoragePricePerEpoch: storagePricePerEpochForDeal,
		ProviderCollateral:   providerCollateral,
	}

	buf, err := cborutil.Dump(&proposal)
	if err != nil {
		return nil, err
	}

	sig, err := sigs.Sign(key.ActSigType(k.Type), k.PrivateKey, buf)
	if err != nil {
		return nil, fmt.Errorf("wallet sign failed: %w", err)
	}

	return &market.ClientDealProposal{
		Proposal:        proposal,
		ClientSignature: *sig,
	}, nil
}

func doRpc(ctx context.Context, s network.Stream, req interface{}, resp interface{}) error {
	errc := make(chan error)
	go func() {
		if err := cborutil.WriteCborRPC(s, req); err != nil {
			errc <- fmt.Errorf("failed to send request: %w", err)
			return
		}

		if err := cborutil.ReadCborRPC(s, resp); err != nil {
			errc <- fmt.Errorf("failed to read response: %w", err)
			return
		}

		errc <- nil
	}()

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func NextPowOf2(n int) int {
	k := 1
	for k < n {
		k = k << 1
	}
	return k
}
