package main

import (
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/boost/storagemarket/types"
	"github.com/filecoin-project/boost/storagemarket/types/dealcheckpoints"
)

const (
	TxInProgress TxDealStatus = "in_progress"
	TxFinished   TxDealStatus = "finished"
	TxFailed     TxDealStatus = "failed"
)

var FlagJsonTx = &cli.BoolFlag{
	Name:  "jsontx",
	Usage: "output results in json format tailored for Telnyx",
	Value: false,
}

type TxDealStatus string

type TxMinerStorageAsk struct {
	Miner              string `json:"miner"`
	PricePerEpoch      string `json:"price_per_epoch"`
	PricePerGB         string `json:"price_per_gb"`
	TotalPrice         string `json:"total_price"`
	VerifiedPricePerGB string `json:"verified_price_per_gb"`
	MaxSize            int64  `json:"max_size_bytes"`
	MinSize            int64  `json:"min_size_bytes"`
}

type TxRetrievalDeal struct {
	OutputPath       string `json:"output_path"`
	Message          string `json:"message"`
	Miner            string `json:"miner"`
	UnsealPrice      string `json:"unseal_price"`
	TransferDuration string `json:"transfer_duration"`
	TransferAvgBPS   int64  `json:"transfer_avg_bps"`
	CreatedAt        string `json:"created_at"`
}

type TxStorageDeal struct {
	DealUUID string `json:"deal_uuid"`
	DealID   string `json:"deal_id"`

	Miner           string `json:"miner"`
	MinerCollateral string `json:"miner_collateral"`

	WalletAddr string `json:"wallet_address"`
	RootCID    string `json:"root_cid"`
	CarPullURL string `json:"car_pull_url"`
	CommP      string `json:"commp"`

	StartEpoch string `json:"start_epoch"`
	EndEpoch   string `json:"end_epoch"`

	Accepted bool   `json:"accepted"`
	Message  string `json:"message"`

	CreatedAt string `json:"created_at"`
}

type TxStorageDealStatus struct {
	DealUUID      string       `json:"deal_uuid"`
	DealCID       string       `json:"deal_cid"`
	DealID        int64        `json:"deal_id"`
	Status        TxDealStatus `json:"deal_status"`
	SealingStatus string       `json:"sealing_status"`
	Label         string       `json:"deal_label"`

	Miner      string `json:"miner"`
	RootCID    string `json:"root_cid"`
	WalletAddr string `json:"wallet_address"`

	Error   string `json:"error"`
	Message string `json:"message"`
}

// statusToDealStatus is based on dealResolver.Message
func statusToDealStatus(resp *types.DealStatusResponse) TxDealStatus {

	switch resp.DealStatus.Status {
	case dealcheckpoints.Accepted.String():
		if resp.IsOffline {
			return TxInProgress
		}
		switch resp.NBytesReceived {
		case 0:
			return TxInProgress
		case 100:
			return TxInProgress
		default:
			return TxInProgress
		}
	case dealcheckpoints.Transferred.String():
		return TxInProgress
	case dealcheckpoints.Published.String():
		return TxInProgress
	case dealcheckpoints.PublishConfirmed.String():
		return TxInProgress
	case dealcheckpoints.AddedPiece.String():
		return TxInProgress
	case dealcheckpoints.IndexedAndAnnounced.String():
		return TxInProgress
	case dealcheckpoints.Complete.String():
		if resp.DealStatus.Error != "" {
			return TxFailed
		}

		return TxFinished
	}

	return TxInProgress
}
