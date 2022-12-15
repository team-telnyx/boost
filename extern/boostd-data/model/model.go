package model

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

// DealInfo is information about a single deal for a given piece
//                      PieceOffset
//                      v
// Sector        [..........................]
// Piece          ......[            ]......
// CAR            ......[      ]............
type DealInfo struct {
	DealUuid    string              `json:"u"`
	IsLegacy    bool                `json:"y"`
	ChainDealID abi.DealID          `json:"i"`
	MinerAddr   address.Address     `json:"m"`
	SectorID    abi.SectorNumber    `json:"s"`
	PieceOffset abi.PaddedPieceSize `json:"o"`
	PieceLength abi.PaddedPieceSize `json:"l"`
	// The size of the CAR file without zero-padding.
	// This value may be zero if the size is unknown.
	CarLength uint64 `json:"c"`

	// If we don't have CarLength, we have to iterate over all offsets, get
	// the largest offset and sum it with length.
}

// Metadata for PieceCid
type Metadata struct {
	IndexedAt time.Time  `json:"i"`
	Deals     []DealInfo `json:"d"`
	Error     string     `json:"e"`
	ErrorType string     `json:"t"`
}

// Record is the information stored in the index for each block in a piece
type Record struct {
	Cid cid.Cid
	OffsetSize
}

type OffsetSize struct {
	// Offset is the offset into the CAR file of the section, where a section
	// is <section size><cid><block data>
	Offset uint64
	// Size is the size of the block data (not the whole section)
	Size uint64
}

func (ofsz *OffsetSize) MarshallBase64() string {
	buf := make([]byte, 2*binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, ofsz.Offset)
	n += binary.PutUvarint(buf[n:], ofsz.Size)
	return base64.StdEncoding.EncodeToString(buf[:n])
}

func (ofsz *OffsetSize) UnmarshallBase64(str string) error {
	buf, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return fmt.Errorf("decoding offset/size from base64 string: %w", err)
	}

	offset, n := binary.Uvarint(buf)
	size, _ := binary.Uvarint(buf[n:])

	ofsz.Offset = offset
	ofsz.Size = size

	return nil
}

// FlaggedPiece is a piece that has been flagged for the user's attention
// (eg because the index is missing)
type FlaggedPiece struct {
	CreatedAt time.Time
	PieceCid  cid.Cid
}
