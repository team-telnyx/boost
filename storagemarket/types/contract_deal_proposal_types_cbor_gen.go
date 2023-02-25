// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package types

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *ContractDealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{173}); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)
	if len("PieceCID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceCID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PieceCID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceCID")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	// t.PieceSize (abi.PaddedPieceSize) (uint64)
	if len("PieceSize") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceSize\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PieceSize"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceSize")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.PieceSize)); err != nil {
		return err
	}

	// t.VerifiedDeal (bool) (bool)
	if len("VerifiedDeal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"VerifiedDeal\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("VerifiedDeal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("VerifiedDeal")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.VerifiedDeal); err != nil {
		return err
	}

	// t.Client (address.Address) (struct)
	if len("Client") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Client\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Client"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Client")); err != nil {
		return err
	}

	if err := t.Client.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Provider (string) (string)
	if len("Provider") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Provider\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Provider"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Provider")); err != nil {
		return err
	}

	if len(t.Provider) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Provider was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Provider))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Provider)); err != nil {
		return err
	}

	// t.Label (market.DealLabel) (struct)
	if len("Label") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Label\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Label"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Label")); err != nil {
		return err
	}

	if err := t.Label.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if len("StartEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StartEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("StartEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StartEpoch")); err != nil {
		return err
	}

	if t.StartEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.StartEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.StartEpoch-1)); err != nil {
			return err
		}
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if len("EndEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"EndEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("EndEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("EndEpoch")); err != nil {
		return err
	}

	if t.EndEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.EndEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.EndEpoch-1)); err != nil {
			return err
		}
	}

	// t.StoragePricePerEpoch (big.Int) (struct)
	if len("StoragePricePerEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StoragePricePerEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("StoragePricePerEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StoragePricePerEpoch")); err != nil {
		return err
	}

	if err := t.StoragePricePerEpoch.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ProviderCollateral (big.Int) (struct)
	if len("ProviderCollateral") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ProviderCollateral\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ProviderCollateral"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ProviderCollateral")); err != nil {
		return err
	}

	if err := t.ProviderCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ClientCollateral (big.Int) (struct)
	if len("ClientCollateral") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ClientCollateral\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ClientCollateral"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ClientCollateral")); err != nil {
		return err
	}

	if err := t.ClientCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Version (string) (string)
	if len("Version") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Version\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Version"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Version")); err != nil {
		return err
	}

	if len(t.Version) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Version was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Version))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Version)); err != nil {
		return err
	}

	// t.Params ([]uint8) (slice)
	if len("Params") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Params\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Params"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Params")); err != nil {
		return err
	}

	if len(t.Params) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Params was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Params))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Params[:]); err != nil {
		return err
	}
	return nil
}

func (t *ContractDealProposal) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ContractDealProposal{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ContractDealProposal: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.PieceCID (cid.Cid) (struct)
		case "PieceCID":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
				}

				t.PieceCID = c

			}
			// t.PieceSize (abi.PaddedPieceSize) (uint64)
		case "PieceSize":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.PieceSize = abi.PaddedPieceSize(extra)

			}
			// t.VerifiedDeal (bool) (bool)
		case "VerifiedDeal":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.VerifiedDeal = false
			case 21:
				t.VerifiedDeal = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.Client (address.Address) (struct)
		case "Client":

			{

				if err := t.Client.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Client: %w", err)
				}

			}
			// t.Provider (string) (string)
		case "Provider":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Provider = string(sval)
			}
			// t.Label (market.DealLabel) (struct)
		case "Label":

			{

				if err := t.Label.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Label: %w", err)
				}

			}
			// t.StartEpoch (abi.ChainEpoch) (int64)
		case "StartEpoch":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.StartEpoch = abi.ChainEpoch(extraI)
			}
			// t.EndEpoch (abi.ChainEpoch) (int64)
		case "EndEpoch":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.EndEpoch = abi.ChainEpoch(extraI)
			}
			// t.StoragePricePerEpoch (big.Int) (struct)
		case "StoragePricePerEpoch":

			{

				if err := t.StoragePricePerEpoch.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.StoragePricePerEpoch: %w", err)
				}

			}
			// t.ProviderCollateral (big.Int) (struct)
		case "ProviderCollateral":

			{

				if err := t.ProviderCollateral.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ProviderCollateral: %w", err)
				}

			}
			// t.ClientCollateral (big.Int) (struct)
		case "ClientCollateral":

			{

				if err := t.ClientCollateral.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ClientCollateral: %w", err)
				}

			}
			// t.Version (string) (string)
		case "Version":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Version = string(sval)
			}
			// t.Params ([]uint8) (slice)
		case "Params":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Params: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Params = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Params[:]); err != nil {
				return err
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ContractParamsVersion1) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.LocationRef (string) (string)
	if len("LocationRef") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"LocationRef\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("LocationRef"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("LocationRef")); err != nil {
		return err
	}

	if len(t.LocationRef) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.LocationRef was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.LocationRef))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.LocationRef)); err != nil {
		return err
	}

	// t.CarSize (uint64) (uint64)
	if len("CarSize") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CarSize\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CarSize"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CarSize")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.CarSize)); err != nil {
		return err
	}

	// t.SkipIpniAnnounce (bool) (bool)
	if len("SkipIpniAnnounce") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SkipIpniAnnounce\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("SkipIpniAnnounce"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SkipIpniAnnounce")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.SkipIpniAnnounce); err != nil {
		return err
	}

	// t.RemoveUnsealedCopy (bool) (bool)
	if len("RemoveUnsealedCopy") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"RemoveUnsealedCopy\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoveUnsealedCopy"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("RemoveUnsealedCopy")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.RemoveUnsealedCopy); err != nil {
		return err
	}
	return nil
}

func (t *ContractParamsVersion1) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ContractParamsVersion1{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ContractParamsVersion1: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.LocationRef (string) (string)
		case "LocationRef":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.LocationRef = string(sval)
			}
			// t.CarSize (uint64) (uint64)
		case "CarSize":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.CarSize = uint64(extra)

			}
			// t.SkipIpniAnnounce (bool) (bool)
		case "SkipIpniAnnounce":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.SkipIpniAnnounce = false
			case 21:
				t.SkipIpniAnnounce = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.RemoveUnsealedCopy (bool) (bool)
		case "RemoveUnsealedCopy":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.RemoveUnsealedCopy = false
			case 21:
				t.RemoveUnsealedCopy = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
