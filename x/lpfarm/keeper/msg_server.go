package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/crescent-network/crescent/v4/x/lpfarm/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// CreatePrivatePlan defines a method to create a new private plan.
func (k msgServer) CreatePrivatePlan(goCtx context.Context, msg *types.MsgCreatePrivatePlan) (*types.MsgCreatePrivatePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	plan, err := k.Keeper.CreatePrivatePlan(
		ctx, creatorAddr, msg.Description, msg.RewardAllocations, msg.StartTime, msg.EndTime)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePrivatePlanResponse{
		PlanId:             plan.Id,
		FarmingPoolAddress: plan.FarmingPoolAddress,
	}, nil
}

// TerminatePrivatePlan defines a method to terminate a private plan.
func (k msgServer) TerminatePrivatePlan(goCtx context.Context, msg *types.MsgTerminatePrivatePlan) (*types.MsgTerminatePrivatePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	plan, found := k.GetPlan(ctx, msg.PlanId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "plan not found: %d", msg.PlanId)
	}
	if !plan.IsPrivate {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot terminate public plan")
	}
	if plan.TerminationAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnauthorized,
			"plan's termination address must be same with the sender's address")
	}

	if err := k.Keeper.TerminatePlan(ctx, plan); err != nil {
		return nil, err
	}

	return &types.MsgTerminatePrivatePlanResponse{}, nil
}

// Farm defines a method for farming coins.
func (k msgServer) Farm(goCtx context.Context, msg *types.MsgFarm) (*types.MsgFarmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	farmerAddr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		return nil, err
	}

	withdrawnRewards, err := k.Keeper.Farm(ctx, farmerAddr, msg.Coin)
	if err != nil {
		return nil, err
	}

	return &types.MsgFarmResponse{
		WithdrawnRewards: withdrawnRewards,
	}, nil
}

// Unfarm defines a method for un-farming coins.
func (k msgServer) Unfarm(goCtx context.Context, msg *types.MsgUnfarm) (*types.MsgUnfarmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	farmerAddr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		return nil, err
	}

	withdrawnRewards, err := k.Keeper.Unfarm(ctx, farmerAddr, msg.Coin)
	if err != nil {
		return nil, err
	}

	return &types.MsgUnfarmResponse{
		WithdrawnRewards: withdrawnRewards,
	}, nil
}

// Harvest defines a method for harvesting farming rewards.
func (k msgServer) Harvest(goCtx context.Context, msg *types.MsgHarvest) (*types.MsgHarvestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	farmerAddr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		return nil, err
	}

	withdrawnRewards, err := k.Keeper.Harvest(ctx, farmerAddr, msg.Denom)
	if err != nil {
		return nil, err
	}

	return &types.MsgHarvestResponse{
		WithdrawnRewards: withdrawnRewards,
	}, nil
}
