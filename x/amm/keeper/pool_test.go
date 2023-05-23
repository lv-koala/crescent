package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/crescent-network/crescent/v5/types"
	"github.com/crescent-network/crescent/v5/x/amm/types"
	exchangetypes "github.com/crescent-network/crescent/v5/x/exchange/types"
)

func (s *KeeperTestSuite) CreateSampleMarketAndPool() (market exchangetypes.Market, pool types.Pool) {
	creatorAddr := s.FundedAccount(0, utils.ParseCoins("1000000ucre,1000000uusd"))
	market = s.CreateMarket(creatorAddr, "ucre", "uusd", true)
	pool = s.CreatePool(creatorAddr, market.Id, sdk.NewDec(5), true)
	return market, pool
}

func (s *KeeperTestSuite) TestPoolOrders() {
	type order struct {
		Price sdk.Dec
		Qty   sdk.Int
	}
	for _, tc := range []struct {
		name         string
		addLiquidity func(pool types.Pool, lpAddr sdk.AccAddress)
		buyOrders    []order
		sellOrders   []order
	}{
		{
			"simple liquidity",
			func(pool types.Pool, lpAddr sdk.AccAddress) {
				s.AddLiquidity(
					lpAddr, lpAddr, pool.Id, utils.ParseDec("4.98"), utils.ParseDec("5.02"),
					utils.ParseCoins("100_000000ucre,500_000000uusd"))
			},
			[]order{
				{utils.ParseDec("4.9950"), sdk.NewInt(25006228)},
				{utils.ParseDec("4.9900"), sdk.NewInt(25043815)},
				{utils.ParseDec("4.9850"), sdk.NewInt(25081497)},
				{utils.ParseDec("4.9800"), sdk.NewInt(25119274)},
			},
			[]order{
				{utils.ParseDec("5.0050"), sdk.NewInt(24956259)},
				{utils.ParseDec("5.0100"), sdk.NewInt(24918890)},
				{utils.ParseDec("5.0150"), sdk.NewInt(24881614)},
				{utils.ParseDec("5.0200"), sdk.NewInt(24844431)},
			},
		},
		{
			"valley",
			func(pool types.Pool, lpAddr sdk.AccAddress) {
				s.AddLiquidity(
					lpAddr, lpAddr, pool.Id, utils.ParseDec("4.96"), utils.ParseDec("4.98"),
					utils.ParseCoins("100_000000ucre,500_000000uusd"))
				s.AddLiquidity(
					lpAddr, lpAddr, pool.Id, utils.ParseDec("5.02"), utils.ParseDec("5.04"),
					utils.ParseCoins("100_000000ucre,500_000000uusd"))
			},
			[]order{
				{utils.ParseDec("4.9750"), sdk.NewInt(25106679)},
				{utils.ParseDec("4.9700"), sdk.NewInt(25144570)},
				{utils.ParseDec("4.9650"), sdk.NewInt(25182556)},
				{utils.ParseDec("4.9600"), sdk.NewInt(25220637)},
			},
			[]order{
				{utils.ParseDec("5.0250"), sdk.NewInt(25055960)},
				{utils.ParseDec("5.0300"), sdk.NewInt(25018591)},
				{utils.ParseDec("5.0350"), sdk.NewInt(24981315)},
				{utils.ParseDec("5.0400"), sdk.NewInt(24944131)},
			},
		},
		{
			"high valley",
			func(pool types.Pool, lpAddr sdk.AccAddress) {
				s.AddLiquidity(
					lpAddr, lpAddr, pool.Id, utils.ParseDec("4.97"), utils.ParseDec("5.03"),
					utils.ParseCoins("100_000000ucre,500_000000uusd"))
				s.AddLiquidity(
					lpAddr, lpAddr, pool.Id, utils.ParseDec("4.98"), utils.ParseDec("4.99"),
					utils.ParseCoins("100_000000ucre,500_000000uusd"))
				s.AddLiquidity(
					lpAddr, lpAddr, pool.Id, utils.ParseDec("5.01"), utils.ParseDec("5.02"),
					utils.ParseCoins("100_000000ucre,500_000000uusd"))
			},
			[]order{
				{utils.ParseDec("4.9950"), sdk.NewInt(16662453)},
				{utils.ParseDec("4.9900"), sdk.NewInt(16687499)},
				{utils.ParseDec("4.9850"), sdk.NewInt(66850484)},
				{utils.ParseDec("4.9800"), sdk.NewInt(66951171)},
				{utils.ParseDec("4.9750"), sdk.NewInt(16763015)},
				{utils.ParseDec("4.9700"), sdk.NewInt(16788313)},
			},
			[]order{
				{utils.ParseDec("5.0050"), sdk.NewInt(16629158)},
				{utils.ParseDec("5.0100"), sdk.NewInt(16604258)},
				{utils.ParseDec("5.0150"), sdk.NewInt(66616807)},
				{utils.ParseDec("5.0200"), sdk.NewInt(66517255)},
				{utils.ParseDec("5.0250"), sdk.NewInt(16529929)},
				{utils.ParseDec("5.0300"), sdk.NewInt(16505276)},
			},
		},
	} {
		s.Run(tc.name, func() {
			s.SetupTest()
			_, pool := s.CreateSampleMarketAndPool()
			lpAddr := s.FundedAccount(1, utils.ParseCoins("10000_000000ucre,10000_000000uusd"))
			tc.addLiquidity(pool, lpAddr)
			var buyOrders, sellOrders []order
			s.App.AMMKeeper.IteratePoolOrders(s.Ctx, pool, true, func(price sdk.Dec, qty sdk.Int, liquidity sdk.Int) (stop bool) {
				buyOrders = append(buyOrders, order{price, qty})
				return false
			})
			s.App.AMMKeeper.IteratePoolOrders(s.Ctx, pool, false, func(price sdk.Dec, qty sdk.Int, liquidity sdk.Int) (stop bool) {
				sellOrders = append(sellOrders, order{price, qty})
				return false
			})
			fmt.Println(buyOrders)
			fmt.Println(sellOrders)
			s.Require().EqualValues(tc.buyOrders, buyOrders)
			s.Require().EqualValues(tc.sellOrders, sellOrders)
		})
	}
}
