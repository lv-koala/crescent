package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (ms msgServer) PricesPrevote(
	goCtx context.Context,
	msg *types.MsgPricesPrevote,
) (*types.MsgPricesPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := ms.Keeper.HandlePrevoteMsg(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgPricesPrevoteResponse{}, nil
}

func (ms msgServer) PricesVote(
	goCtx context.Context,
	msg *types.MsgPricesVote,
) (*types.MsgPricesVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := ms.Keeper.HandleVoteMsg(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgPricesVoteResponse{}, nil
}

func (ms msgServer) AddNewTicker(
	goCtx context.Context,
	msg *types.MsgAddNewTicker,
) (*types.MsgAddNewTickerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := ms.Keeper.HandleAddNewTickerMsg(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgAddNewTickerResponse{}, nil
}

func (ms msgServer) LsvFeederUpdate(
	goCtx context.Context,
	msg *types.MsgLsvFeederUpdate,
) (*types.MsgLsvFeederUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := ms.Keeper.HandleLsvFeederUpdateMsg(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgLsvFeederUpdateResponse{}, nil
}
