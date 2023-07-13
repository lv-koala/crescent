package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

var _ types.QueryServer = Querier{}

// Querier implements a QueryServer for the x/oracle module.
type Querier struct {
	Keeper
}

// NewQuerier returns an implementation of the oracle QueryServer interface
// for the provided Keeper.
func NewQuerier(keeper Keeper) types.QueryServer {
	return &Querier{Keeper: keeper}
}

// Params queries params of x/oracle module.
func (q Querier) Params(
	goCtx context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Pricess queries exchange rates of all denoms, or, if specified, returns
// a single denom.
func (q Querier) PriceOracle(
	goCtx context.Context,
	req *types.QueryPriceOracle,
) (*types.QueryPriceOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var op types.OraclePrices

	q.IteratePrices(ctx, func(pr types.OraclePrice) (stop bool) {
		op = append(op, pr)
		return false
	})

	return &types.QueryPriceOracleResponse{PriceOracle: op}, nil
}

// FeederDelegation queries the account address to which the validator operator
// delegated oracle vote rights.
func (q Querier) Prevotes(
	goCtx context.Context,
	req *types.QueryPrevotes,
) (*types.QueryPrevotesResponse, error) {

	return &types.QueryPrevotesResponse{}, nil
}

// MissCounter queries oracle miss counter of a validator.
func (q Querier) MissCounter(
	goCtx context.Context,
	req *types.QueryMissCounter,
) (*types.QueryMissCounterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	/*
		valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		ctx := sdk.UnwrapSDKContext(goCtx)

		return &types.QueryMissCounterResponse{
			MissCounter: q.GetMissCounter(ctx, valAddr),
		}, nil
	*/
	return &types.QueryMissCounterResponse{}, nil

}

// AggregatePrevotes queries aggregate prevotes of all validators
func (q Querier) Votes(
	goCtx context.Context,
	req *types.QueryVotes,
) (*types.QueryVotesResponse, error) {

	return &types.QueryVotesResponse{}, nil
}
