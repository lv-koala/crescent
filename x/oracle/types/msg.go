package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

var (
	_ sdk.Msg = &MsgPricesPrevote{}
	_ sdk.Msg = &MsgPricesVote{}
	_ sdk.Msg = &MsgLsvFeederUpdate{}
	_ sdk.Msg = &MsgAddNewTicker{}
)

// Message types for the module
const (
	TypeMsgPrevote         = "prices_prevote"
	TypeMsgVote            = "prices_vote"
	TypeMsgLsvFeederUpdate = "lsv_feeder_update"
	TypeMsgAddNewSymbol    = "add_new_ticker"
)

func NewMsgPricesPrevote(
	hash string, feeder sdk.AccAddress) *MsgPricesPrevote {
	return &MsgPricesPrevote{
		Hash:   hash,
		Feeder: feeder.String(),
	}
}

func (msg MsgPricesPrevote) Route() string { return RouterKey }
func (msg MsgPricesPrevote) Type() string  { return TypeMsgPrevote }

func (msg MsgPricesPrevote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPricesPrevote) GetSigners() []sdk.AccAddress {
	//feeder == msgSigner
	feeder, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{feeder}
}

func (msg MsgPricesPrevote) ValidateBasic() error {

	fmt.Printf("PREVOTE validate %v\n", msg.Feeder)
	fmt.Printf("%v\n", len(msg.Hash))
	fmt.Printf("%v\n", msg.Hash)
	_, err := AggregateVoteHashFromHex(msg.Hash)
	if err != nil {
		fmt.Println(err)
		return sdkerrors.Wrapf(ErrInvalidHash, "Invalid vote hash (%s)", err)
	}

	// HEX encoding doubles the hash length
	if len(msg.Hash) != tmhash.TruncatedSize*2 {
		fmt.Printf("no hex")
		return ErrInvalidHashLength
	}

	if _, err := sdk.AccAddressFromBech32(msg.Feeder); err != nil {
		fmt.Printf("feeder")
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feeder address: %v", err)
	}
	return nil
}

// VOTE
func NewMsgPricesVote(
	salt string, pricesTuplesStr string, feeder sdk.AccAddress) *MsgPricesVote {
	return &MsgPricesVote{
		Salt:        salt,
		Feeder:      feeder.String(),
		PriceTuples: pricesTuplesStr,
	}
}

func (msg MsgPricesVote) Route() string { return RouterKey }
func (msg MsgPricesVote) Type() string  { return TypeMsgVote }

func (msg MsgPricesVote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPricesVote) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgPricesVote) ValidateBasic() error {
	fmt.Printf("VOTE validate %v\n", msg.Feeder)

	if _, err := sdk.AccAddressFromBech32(msg.Feeder); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	return nil
}

// LSV Feeder update
func NewMsgLsvFeederUpdate(
	lsvAddr sdk.ValAddress, feeder sdk.AccAddress) *MsgLsvFeederUpdate {
	return &MsgLsvFeederUpdate{
		LsvAccAddr: lsvAddr.String(),
		Feeder:     feeder.String(),
	}
}

func (msg MsgLsvFeederUpdate) Route() string { return RouterKey }
func (msg MsgLsvFeederUpdate) Type() string  { return TypeMsgVote }

func (msg MsgLsvFeederUpdate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgLsvFeederUpdate) GetSigners() []sdk.AccAddress {
	//TODO: validator use a same key as a AccAddress
	addr, err := sdk.AccAddressFromBech32(msg.LsvAccAddr)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (msg MsgLsvFeederUpdate) ValidateBasic() error {
	fmt.Printf("LsvFeederUpdate validate %v\n", msg.Feeder)

	if _, err := sdk.AccAddressFromBech32(msg.Feeder); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	return nil
}

// Add new symbol
func NewMsgAddNewSymbol(
	delegateAddr string, ticker string) *MsgAddNewTicker {
	return &MsgAddNewTicker{
		DelegateAddress: delegateAddr,
		Ticker:          ticker,
	}
}

func (msg MsgAddNewTicker) Route() string { return RouterKey }
func (msg MsgAddNewTicker) Type() string  { return TypeMsgVote }

func (msg MsgAddNewTicker) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgAddNewTicker) GetSigners() []sdk.AccAddress {
	//TODO: check signer is in lsv list
	addr, err := sdk.AccAddressFromBech32(msg.DelegateAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgAddNewTicker) ValidateBasic() error {
	fmt.Printf("AddNewTicker validate %v\n", msg.DelegateAddress)

	if _, err := sdk.AccAddressFromBech32(msg.DelegateAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	return nil
}
