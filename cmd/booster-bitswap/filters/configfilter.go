package filters

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/dustin/go-humanize"
	"github.com/filecoin-project/boost/cmd/booster-bitswap/requestcounter"
	"github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

// PeerListType is either an allow list or a deny list
type PeerListType string

// AllowList is a peer list where only the specified peers are allowed to serve retrievals
const AllowList PeerListType = "allowlist"

// DenyList is a peer list where the specified peers cannot serve retrievals, but all others can
const DenyList PeerListType = "denylist"

// BandwidthMeasure provides an up to date measurement of the current bytes per second transferred
type BandwidthMeasure interface {
	AvgBytesPerSecond() uint64
}

type RequestCounter interface {
	StateForPeer(p peer.ID) requestcounter.ServerState
}

type remoteConfig struct {
	peerListType                   PeerListType
	peerList                       map[peer.ID]struct{}
	underMaintenance               bool
	maxSimultaneousRequests        uint64
	maxSimultaneousRequestsPerPeer uint64
	maxBandwidth                   uint64
}

// ConfigFilter manages filtering based on a remotely fetched retrieval configuration
type ConfigFilter struct {
	remoteConfigLk   sync.RWMutex
	bandwidthMeasure BandwidthMeasure
	requestCounter   RequestCounter
	remoteConfig     remoteConfig
}

// NewConfigFilter constructs a new peer filter
func NewConfigFilter(bandwidthMeasure BandwidthMeasure, requestCounter RequestCounter) *ConfigFilter {
	return &ConfigFilter{
		bandwidthMeasure: bandwidthMeasure,
		requestCounter:   requestCounter,
		remoteConfig: remoteConfig{
			peerListType:                   DenyList,
			peerList:                       make(map[peer.ID]struct{}),
			underMaintenance:               false,
			maxSimultaneousRequests:        0,
			maxSimultaneousRequestsPerPeer: 0,
			maxBandwidth:                   0,
		},
	}
}

// FulfillRequest checks if a given peer is in the allow/deny list and decides
// whether to fulfill the request
func (cf *ConfigFilter) FulfillRequest(p peer.ID, c cid.Cid) (bool, error) {
	cf.remoteConfigLk.RLock()
	defer cf.remoteConfigLk.RUnlock()
	// don't fulfill requests under maintainence
	if cf.remoteConfig.underMaintenance {
		return false, nil
	}
	// don't fulfill requests for peers on deny list or not on an allowlist
	_, has := cf.remoteConfig.peerList[p]
	if (cf.remoteConfig.peerListType == DenyList) == has {
		return false, nil
	}
	// don't fulfill requests when over maxbandwidth
	if cf.remoteConfig.maxBandwidth > 0 && cf.bandwidthMeasure.AvgBytesPerSecond() > cf.remoteConfig.maxBandwidth {
		return false, nil
	}
	s := cf.requestCounter.StateForPeer(p)
	// don't fulfill requests when there are too many simultaneous requests over all
	if cf.remoteConfig.maxSimultaneousRequests > 0 && s.TotalRequestsInProgress >= cf.remoteConfig.maxSimultaneousRequests {
		return false, nil
	}
	// don't fulfill requests when there are too many simultaneous requests for this peer
	if cf.remoteConfig.maxSimultaneousRequestsPerPeer > 0 && s.RequestsInProgressForPeer >= cf.remoteConfig.maxSimultaneousRequestsPerPeer {
		return false, nil
	}
	// all filters passed, fulfill
	return true, nil
}

// parse a response from the peer filter endpoint to get a new set of allowed/denied peers
// and other configs
func (cf *ConfigFilter) parseRemoteConfig(response io.Reader) (remoteConfig, error) {
	type allowDenyList struct {
		Type    string   `json:"Type"`
		PeerIDs []string `json:"PeerIDs"`
	}

	type bitswapLimits struct {
		SimultaneousRequests        uint64 `json:"SimultaneousRequests"`
		SimultaneousRequestsPerPeer uint64 `json:"SimultaneousRequestsPerPeer"`
		MaxBandwidth                string `json:"MaxBandwidth"`
	}

	type storageProviderLimits struct {
		Bitswap bitswapLimits `json:"Bitswap"`
	}
	type responseType struct {
		UnderMaintenance      bool                  `json:"UnderMaintenance"`
		AllowDenyList         allowDenyList         `json:"AllowDenyList"`
		StorageProviderLimits storageProviderLimits `json:"StorageProviderLimits"`
	}

	jsonResponse := json.NewDecoder(response)
	// initialize a json decoder
	var decodedResponse responseType
	err := jsonResponse.Decode(&decodedResponse)
	// read open bracket
	if err != nil {
		return remoteConfig{}, fmt.Errorf("parsing response: %w", err)
	}

	peerListType := DenyList
	if decodedResponse.AllowDenyList.Type != "" {
		if decodedResponse.AllowDenyList.Type != string(DenyList) && decodedResponse.AllowDenyList.Type != string(AllowList) {
			return remoteConfig{}, fmt.Errorf("parsing response: 'Type' must be either '%s' or '%s'", AllowList, DenyList)
		}
		peerListType = PeerListType(decodedResponse.AllowDenyList.Type)
	}

	peerList := make(map[peer.ID]struct{}, len(decodedResponse.AllowDenyList.PeerIDs))
	// while the array contains values
	for _, peerString := range decodedResponse.AllowDenyList.PeerIDs {
		peerID, err := peer.Decode(peerString)
		if err != nil {
			return remoteConfig{}, fmt.Errorf("parsing response: %w", err)
		}
		peerList[peerID] = struct{}{}
	}

	maxBandwidth := uint64(0)
	if decodedResponse.StorageProviderLimits.Bitswap.MaxBandwidth != "" {
		maxBandwidth, err = humanize.ParseBytes(decodedResponse.StorageProviderLimits.Bitswap.MaxBandwidth)
		if err != nil {
			return remoteConfig{}, fmt.Errorf("parsing response: parsing 'MaxBandwidth': %w", err)
		}
	}
	return remoteConfig{
		underMaintenance:               decodedResponse.UnderMaintenance,
		maxSimultaneousRequests:        decodedResponse.StorageProviderLimits.Bitswap.SimultaneousRequests,
		maxSimultaneousRequestsPerPeer: decodedResponse.StorageProviderLimits.Bitswap.SimultaneousRequestsPerPeer,
		maxBandwidth:                   maxBandwidth,
		peerListType:                   peerListType,
		peerList:                       peerList,
	}, nil
}

// ParseUpdate parses and updates the Peer filter list based on an endpoint response
func (cf *ConfigFilter) ParseUpdate(stream io.Reader) error {
	remoteConfig, err := cf.parseRemoteConfig(stream)
	if err != nil {
		return err
	}
	cf.remoteConfigLk.Lock()
	cf.remoteConfig = remoteConfig
	cf.remoteConfigLk.Unlock()
	return nil
}
