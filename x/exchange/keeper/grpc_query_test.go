package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/crescent-network/crescent/v5/types"
	ammtypes "github.com/crescent-network/crescent/v5/x/amm/types"
	"github.com/crescent-network/crescent/v5/x/exchange/keeper"
	"github.com/crescent-network/crescent/v5/x/exchange/types"
)

func (s *KeeperTestSuite) TestQueryBestSwapExactInRoutes() {
	creatorAddr := utils.TestAddress(1)
	s.FundAccount(creatorAddr, utils.ParseCoins("100000_000000ucre,100000_000000uatom,100000_000000uusd"))

	market1 := s.CreateSpotMarket(utils.TestAddress(0), "ucre", "uusd", true)
	market2 := s.CreateSpotMarket(utils.TestAddress(0), "uatom", "ucre", true)
	market3 := s.CreateSpotMarket(utils.TestAddress(0), "uatom", "uusd", true)

	pool1 := s.CreatePool(creatorAddr, market1.Id, ammtypes.DefaultDefaultTickSpacing, utils.ParseDec("9.7"), true)
	s.AddLiquidity(creatorAddr, pool1.Id, utils.ParseDec("9.5"), utils.ParseDec("10"),
		sdk.NewInt(10000_000000), sdk.NewInt(10000_000000), sdk.OneInt(), sdk.OneInt())
	pool2 := s.CreatePool(creatorAddr, market2.Id, ammtypes.DefaultDefaultTickSpacing, utils.ParseDec("1.04"), true)
	s.AddLiquidity(creatorAddr, pool2.Id, utils.ParseDec("1"), utils.ParseDec("1.2"),
		sdk.NewInt(10000_000000), sdk.NewInt(10000_000000), sdk.OneInt(), sdk.OneInt())
	pool3 := s.CreatePool(creatorAddr, market3.Id, ammtypes.DefaultDefaultTickSpacing, utils.ParseDec("10.3"), true)
	s.AddLiquidity(creatorAddr, pool3.Id, utils.ParseDec("9.7"), utils.ParseDec("11"),
		sdk.NewInt(10000_000000), sdk.NewInt(10000_000000), sdk.OneInt(), sdk.OneInt())

	querier := keeper.Querier{Keeper: s.App.ExchangeKeeper}
	resp, err := querier.BestSwapExactInRoutes(sdk.WrapSDKContext(s.Ctx), &types.QueryBestSwapExactInRoutesRequest{
		Input:       utils.ParseCoin("100_000000ucre"),
		OutputDenom: "uusd",
	})
	s.Require().NoError(err)

	s.Require().EqualValues([]uint64{2, 3}, resp.Routes)
}
