// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package datacap

import (
	"fmt"
	"io"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"

	address "github.com/jimpick/go-address"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

var lengthBufState = []byte{130}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}

	// t.Governor (address.Address) (struct)
	if err := t.Governor.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Token (datacap.TokenState) (struct)
	if err := t.Token.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) error {
	*t = State{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Governor (address.Address) (struct)

	{

		if err := t.Governor.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Governor: %w", err)
		}

	}
	// t.Token (datacap.TokenState) (struct)

	{

		if err := t.Token.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Token: %w", err)
		}

	}
	return nil
}

var lengthBufTokenState = []byte{132}

func (t *TokenState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufTokenState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Supply (big.Int) (struct)
	if err := t.Supply.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Balances (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Balances); err != nil {
		return xerrors.Errorf("failed to write cid field t.Balances: %w", err)
	}

	// t.Allowances (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Allowances); err != nil {
		return xerrors.Errorf("failed to write cid field t.Allowances: %w", err)
	}

	// t.HamtBitWidth (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.HamtBitWidth)); err != nil {
		return err
	}

	return nil
}

func (t *TokenState) UnmarshalCBOR(r io.Reader) error {
	*t = TokenState{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Supply (big.Int) (struct)

	{

		if err := t.Supply.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Supply: %w", err)
		}

	}
	// t.Balances (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Balances: %w", err)
		}

		t.Balances = c

	}
	// t.Allowances (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Allowances: %w", err)
		}

		t.Allowances = c

	}
	// t.HamtBitWidth (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.HamtBitWidth = uint64(extra)

	}
	return nil
}

var lengthBufMintParams = []byte{131}

func (t *MintParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufMintParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.To (address.Address) (struct)
	if err := t.To.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Operators ([]address.Address) (slice)
	if len(t.Operators) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Operators was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Operators))); err != nil {
		return err
	}
	for _, v := range t.Operators {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *MintParams) UnmarshalCBOR(r io.Reader) error {
	*t = MintParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.To (address.Address) (struct)

	{

		if err := t.To.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.To: %w", err)
		}

	}
	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	// t.Operators ([]address.Address) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Operators: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Operators = make([]address.Address, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v address.Address
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Operators[i] = v
	}

	return nil
}

var lengthBufMintReturn = []byte{131}

func (t *MintReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufMintReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Balance (big.Int) (struct)
	if err := t.Balance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Supply (big.Int) (struct)
	if err := t.Supply.MarshalCBOR(w); err != nil {
		return err
	}

	// t.RecipientData ([]uint8) (slice)
	if len(t.RecipientData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.RecipientData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.RecipientData))); err != nil {
		return err
	}

	if _, err := w.Write(t.RecipientData[:]); err != nil {
		return err
	}
	return nil
}

func (t *MintReturn) UnmarshalCBOR(r io.Reader) error {
	*t = MintReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Balance (big.Int) (struct)

	{

		if err := t.Balance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Balance: %w", err)
		}

	}
	// t.Supply (big.Int) (struct)

	{

		if err := t.Supply.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Supply: %w", err)
		}

	}
	// t.RecipientData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.RecipientData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.RecipientData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.RecipientData[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufDestroyParams = []byte{130}

func (t *DestroyParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDestroyParams); err != nil {
		return err
	}

	// t.Owner (address.Address) (struct)
	if err := t.Owner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DestroyParams) UnmarshalCBOR(r io.Reader) error {
	*t = DestroyParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Owner (address.Address) (struct)

	{

		if err := t.Owner.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Owner: %w", err)
		}

	}
	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	return nil
}

var lengthBufTransferParams = []byte{131}

func (t *TransferParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufTransferParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.To (address.Address) (struct)
	if err := t.To.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}

	// t.OperatorData ([]uint8) (slice)
	if len(t.OperatorData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.OperatorData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.OperatorData))); err != nil {
		return err
	}

	if _, err := w.Write(t.OperatorData[:]); err != nil {
		return err
	}
	return nil
}

func (t *TransferParams) UnmarshalCBOR(r io.Reader) error {
	*t = TransferParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.To (address.Address) (struct)

	{

		if err := t.To.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.To: %w", err)
		}

	}
	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	// t.OperatorData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.OperatorData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.OperatorData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.OperatorData[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufTransferReturn = []byte{131}

func (t *TransferReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufTransferReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.FromBalance (big.Int) (struct)
	if err := t.FromBalance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ToBalance (big.Int) (struct)
	if err := t.ToBalance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.RecipientData ([]uint8) (slice)
	if len(t.RecipientData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.RecipientData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.RecipientData))); err != nil {
		return err
	}

	if _, err := w.Write(t.RecipientData[:]); err != nil {
		return err
	}
	return nil
}

func (t *TransferReturn) UnmarshalCBOR(r io.Reader) error {
	*t = TransferReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.FromBalance (big.Int) (struct)

	{

		if err := t.FromBalance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.FromBalance: %w", err)
		}

	}
	// t.ToBalance (big.Int) (struct)

	{

		if err := t.ToBalance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ToBalance: %w", err)
		}

	}
	// t.RecipientData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.RecipientData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.RecipientData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.RecipientData[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufTransferFromParams = []byte{132}

func (t *TransferFromParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufTransferFromParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.From (address.Address) (struct)
	if err := t.From.MarshalCBOR(w); err != nil {
		return err
	}

	// t.To (address.Address) (struct)
	if err := t.To.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}

	// t.OperatorData ([]uint8) (slice)
	if len(t.OperatorData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.OperatorData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.OperatorData))); err != nil {
		return err
	}

	if _, err := w.Write(t.OperatorData[:]); err != nil {
		return err
	}
	return nil
}

func (t *TransferFromParams) UnmarshalCBOR(r io.Reader) error {
	*t = TransferFromParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.From (address.Address) (struct)

	{

		if err := t.From.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.From: %w", err)
		}

	}
	// t.To (address.Address) (struct)

	{

		if err := t.To.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.To: %w", err)
		}

	}
	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	// t.OperatorData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.OperatorData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.OperatorData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.OperatorData[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufTransferFromReturn = []byte{132}

func (t *TransferFromReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufTransferFromReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.FromBalance (big.Int) (struct)
	if err := t.FromBalance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ToBalance (big.Int) (struct)
	if err := t.ToBalance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Allowance (big.Int) (struct)
	if err := t.Allowance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.RecipientData ([]uint8) (slice)
	if len(t.RecipientData) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.RecipientData was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.RecipientData))); err != nil {
		return err
	}

	if _, err := w.Write(t.RecipientData[:]); err != nil {
		return err
	}
	return nil
}

func (t *TransferFromReturn) UnmarshalCBOR(r io.Reader) error {
	*t = TransferFromReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.FromBalance (big.Int) (struct)

	{

		if err := t.FromBalance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.FromBalance: %w", err)
		}

	}
	// t.ToBalance (big.Int) (struct)

	{

		if err := t.ToBalance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ToBalance: %w", err)
		}

	}
	// t.Allowance (big.Int) (struct)

	{

		if err := t.Allowance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Allowance: %w", err)
		}

	}
	// t.RecipientData ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.RecipientData: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.RecipientData = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.RecipientData[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufIncreaseAllowanceParams = []byte{130}

func (t *IncreaseAllowanceParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufIncreaseAllowanceParams); err != nil {
		return err
	}

	// t.Operator (address.Address) (struct)
	if err := t.Operator.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Increase (big.Int) (struct)
	if err := t.Increase.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *IncreaseAllowanceParams) UnmarshalCBOR(r io.Reader) error {
	*t = IncreaseAllowanceParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Operator (address.Address) (struct)

	{

		if err := t.Operator.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Operator: %w", err)
		}

	}
	// t.Increase (big.Int) (struct)

	{

		if err := t.Increase.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Increase: %w", err)
		}

	}
	return nil
}

var lengthBufDecreaseAllowanceParams = []byte{130}

func (t *DecreaseAllowanceParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDecreaseAllowanceParams); err != nil {
		return err
	}

	// t.Operator (address.Address) (struct)
	if err := t.Operator.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Decrease (big.Int) (struct)
	if err := t.Decrease.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DecreaseAllowanceParams) UnmarshalCBOR(r io.Reader) error {
	*t = DecreaseAllowanceParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Operator (address.Address) (struct)

	{

		if err := t.Operator.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Operator: %w", err)
		}

	}
	// t.Decrease (big.Int) (struct)

	{

		if err := t.Decrease.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Decrease: %w", err)
		}

	}
	return nil
}

var lengthBufRevokeAllowanceParams = []byte{129}

func (t *RevokeAllowanceParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufRevokeAllowanceParams); err != nil {
		return err
	}

	// t.Operator (address.Address) (struct)
	if err := t.Operator.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *RevokeAllowanceParams) UnmarshalCBOR(r io.Reader) error {
	*t = RevokeAllowanceParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Operator (address.Address) (struct)

	{

		if err := t.Operator.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Operator: %w", err)
		}

	}
	return nil
}

var lengthBufGetAllowanceParams = []byte{130}

func (t *GetAllowanceParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufGetAllowanceParams); err != nil {
		return err
	}

	// t.Owner (address.Address) (struct)
	if err := t.Owner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Operator (address.Address) (struct)
	if err := t.Operator.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *GetAllowanceParams) UnmarshalCBOR(r io.Reader) error {
	*t = GetAllowanceParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Owner (address.Address) (struct)

	{

		if err := t.Owner.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Owner: %w", err)
		}

	}
	// t.Operator (address.Address) (struct)

	{

		if err := t.Operator.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Operator: %w", err)
		}

	}
	return nil
}

var lengthBufBurnParams = []byte{129}

func (t *BurnParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufBurnParams); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *BurnParams) UnmarshalCBOR(r io.Reader) error {
	*t = BurnParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	return nil
}

var lengthBufBurnReturn = []byte{129}

func (t *BurnReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufBurnReturn); err != nil {
		return err
	}

	// t.Balance (big.Int) (struct)
	if err := t.Balance.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *BurnReturn) UnmarshalCBOR(r io.Reader) error {
	*t = BurnReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Balance (big.Int) (struct)

	{

		if err := t.Balance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Balance: %w", err)
		}

	}
	return nil
}

var lengthBufBurnFromParams = []byte{130}

func (t *BurnFromParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufBurnFromParams); err != nil {
		return err
	}

	// t.Owner (address.Address) (struct)
	if err := t.Owner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *BurnFromParams) UnmarshalCBOR(r io.Reader) error {
	*t = BurnFromParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Owner (address.Address) (struct)

	{

		if err := t.Owner.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Owner: %w", err)
		}

	}
	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	return nil
}

var lengthBufBurnFromReturn = []byte{130}

func (t *BurnFromReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufBurnFromReturn); err != nil {
		return err
	}

	// t.Balance (big.Int) (struct)
	if err := t.Balance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Allowance (big.Int) (struct)
	if err := t.Allowance.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *BurnFromReturn) UnmarshalCBOR(r io.Reader) error {
	*t = BurnFromReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Balance (big.Int) (struct)

	{

		if err := t.Balance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Balance: %w", err)
		}

	}
	// t.Allowance (big.Int) (struct)

	{

		if err := t.Allowance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Allowance: %w", err)
		}

	}
	return nil
}
