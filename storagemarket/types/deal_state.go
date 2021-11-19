package types

import (
	"time"

	"github.com/filecoin-project/boost/storagemarket/types/dealcheckpoints"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

// ProviderDealState is the local state tracked for a deal by the StorageProvider.
type ProviderDealState struct {
	// DealUuid is an unique uuid generated by client for the deal.
	DealUuid uuid.UUID
	// CreatedAt is the time at which the deal was stored
	CreatedAt time.Time
	// ClientDealProposal is the deal proposal sent by the client.
	ClientDealProposal market.ClientDealProposal

	// SelfPeerID is the Storage Provider's libp2p Peer ID.
	SelfPeerID peer.ID
	// ClientPeerID is the Clients libp2p Peer ID.
	ClientPeerID peer.ID

	// DealDataRoot is the root of the IPLD DAG that the client wants to store.
	DealDataRoot cid.Cid

	// data-transfer
	// InboundCARPath is the file-path where the storage provider will persist the CAR file sent by the client.
	InboundFilePath string
	// TransferURL is the URL sent by the client where the Storage Provider can fetch the CAR file from.
	//TransferURL    string
	TransferType   string
	TransferParams []byte

	// Chain Vars
	ChainDealID abi.DealID
	PublishCID  *cid.Cid

	// sector packing info
	SectorID abi.SectorNumber
	Offset   abi.PaddedPieceSize
	Length   abi.PaddedPieceSize

	// deal checkpoint in DB.
	Checkpoint dealcheckpoints.Checkpoint
	// set if there's an error
	Err string

	// NBytesReceived is the number of bytes Received for this deal
	NBytesReceived int64
}