package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/crescent-network/crescent/v5/x/oracle/keeper"
	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

// NewHandler returns a handler for "oracle" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgPricesPrevote:
			res, err := msgServer.PricesPrevote(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPricesVote:
			res, err := msgServer.PricesVote(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgLsvFeederUpdate:
			res, err := msgServer.LsvFeederUpdate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddNewTicker:
			res, err := msgServer.AddNewTicker(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized oracle message type: %T", msg)
		}
	}
}
