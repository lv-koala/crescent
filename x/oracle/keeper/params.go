package keeper

import (
	"github.com/crescent-network/crescent/v5/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VotePeriod returns the number of blocks during which voting takes place.
func (k Keeper) VoteThreshold(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, types.KeyVoteThreshold, &res)
	return
}

// VoteThreshold returns the minimum percentage of votes that must be received for a ballot to pass.
func (k Keeper) RewardDistributionRate(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, types.KeyRewardRate, &res)
	return
}

// RewardBand returns the ratio of allowable exchange rate error that a validator can be rewared
func (k Keeper) RewardDistributionPeriod(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, types.KeyRewardDistributionPeriod, &res)
	return
}

// RewardDistributionWindow returns the number of vote periods during which seigiornage reward comes in and then is distributed.
func (k Keeper) WhitelistFeeders(ctx sdk.Context) (res []string) {
	k.paramSpace.Get(ctx, types.KeyWhitelistFeeders, &res)
	return
}

func (k Keeper) AcceptTickers(ctx sdk.Context) (res []string) {
	k.paramSpace.Get(ctx, types.KeyAcceptTickers, &res)
	return
}

func (k Keeper) DelegateAddr(ctx sdk.Context) (res []string) {
	k.paramSpace.Get(ctx, types.KeyDelegateAccounts, &res)
	return
}

func (k Keeper) MaxAcceptLatencySeconds(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, types.KeyMaxAcceptLatencySeconds, &res)
	return
}

// Set
func (k Keeper) SetVoteThreshold(ctx sdk.Context, voteThreshold uint64) {
	k.paramSpace.Set(ctx, types.KeyVoteThreshold, voteThreshold)
}

func (k Keeper) SetRewardDistributionRate(ctx sdk.Context, rewardDistributionRate uint64) {
	k.paramSpace.Set(ctx, types.KeyRewardRate, rewardDistributionRate)
}

func (k Keeper) SetRewardDistributionPeriod(ctx sdk.Context, rewardDistributionPeriod uint64) {
	k.paramSpace.Set(ctx, types.KeyRewardDistributionPeriod, rewardDistributionPeriod)
}

func (k Keeper) SetWhitelistFeeders(ctx sdk.Context, whitelist []string) {
	k.paramSpace.Set(ctx, types.KeyWhitelistFeeders, whitelist)
}

func (k Keeper) SetAcceptSymbols(ctx sdk.Context, acceptSymbols []string) {
	k.paramSpace.Set(ctx, types.KeyAcceptTickers, acceptSymbols)
}

func (k Keeper) SetDelegateAddr(ctx sdk.Context, deleAddr []string) {
	k.paramSpace.Set(ctx, types.KeyDelegateAccounts, deleAddr)
}

func (k Keeper) SetMaxAcceptLatencySeconds(ctx sdk.Context, maxAcceptLatencySeconds uint64) {
	k.paramSpace.Set(ctx, types.KeyMaxAcceptLatencySeconds, maxAcceptLatencySeconds)
}

// GetParams returns the total set of oracle parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of oracle parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
