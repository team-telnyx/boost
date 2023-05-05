// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package types

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *LegsVoucher) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{161}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Head (cid.Cid) (struct)
	if len("Head") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Head\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Head"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Head")); err != nil {
		return err
	}

	if t.Head == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.Head); err != nil {
			return xerrors.Errorf("failed to write cid field t.Head: %w", err)
		}
	}

	return nil
}

func (t *LegsVoucher) UnmarshalCBOR(r io.Reader) error {
	*t = LegsVoucher{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("LegsVoucher: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Head (cid.Cid) (struct)
		case "Head":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.Head: %w", err)
					}

					t.Head = &c
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *LegsVoucherResult) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{161}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Code (uint64) (uint64)
	if len("Code") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Code\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Code"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Code")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Code)); err != nil {
		return err
	}

	return nil
}

func (t *LegsVoucherResult) UnmarshalCBOR(r io.Reader) error {
	*t = LegsVoucherResult{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("LegsVoucherResult: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Code (uint64) (uint64)
		case "Code":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Code = uint64(extra)

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
