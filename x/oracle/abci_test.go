package oracle_test

import (
	"fmt"

	"github.com/crescent-network/crescent/v5/types"
	"github.com/crescent-network/crescent/v5/x/oracle"

	_ "github.com/stretchr/testify/suite"
)

// TEST prevote, vote with 1 symbol
func (suite *ModuleTestSuite) TestMidBlockerPriceTest() {
	suite.SetupTest()

	//PREVOTE BLOCK SET
	prevoteHeight := int64(1)
	t := types.ParseTime("2023-05-01T00:00:00Z") // 1682899200
	suite.ctx = suite.ctx.WithBlockTime(t)
	suite.ctx = suite.ctx.WithBlockHeight(prevoteHeight)

	prevoteBlockHight := suite.ctx.BlockHeight()

	//PREVOTE
	suite.Prevote(suite.addrs[0], "abcdef", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5")
	suite.Prevote(suite.addrs[1], "123456", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5")
	suite.Prevote(suite.addrs[2], "ababab", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5")
	suite.Prevote(suite.addrs[3], "dedede", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5")
	suite.Prevote(suite.addrs[4], "aaaaaa", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5")
	suite.Prevote(suite.addrs[5], "bbbbbb", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5")

	//VOTE BLOCK SET
	t2 := types.ParseTime("2023-05-01T00:00:05Z") // 1682899205	// +5
	suite.ctx = suite.ctx.WithBlockTime(t2)
	suite.ctx = suite.ctx.WithBlockHeight(prevoteHeight + 1)

	//VOTE
	suite.Vote(suite.addrs[0], "abcdef", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5", prevoteBlockHight)
	suite.Vote(suite.addrs[1], "123456", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5", prevoteBlockHight)
	suite.Vote(suite.addrs[2], "ababab", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5", prevoteBlockHight)
	suite.Vote(suite.addrs[3], "dedede", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5", prevoteBlockHight)
	suite.Vote(suite.addrs[4], "aaaaaa", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5", prevoteBlockHight)
	suite.Vote(suite.addrs[5], "bbbbbb", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5", prevoteBlockHight)

	oracle.MidBlocker(suite.ctx, suite.keeper)

	//TODO: delete. debug
	price := suite.querier.Keeper.AllOraclePrices(suite.ctx)
	fmt.Println(price)

	suite.Require().Equal(1, len(price))
	suite.Require().Equal("ETH", price[0].Ticker)
	suite.Require().Equal("1808.150000000000000000", price[0].UsdPrice.String()) // 18 precision

}

func (suite *ModuleTestSuite) TestMidBlockerPriceTest2() {
	suite.SetupTest()

	//PREVOTE BLOCK SET
	prevoteHeight := int64(1)
	t := types.ParseTime("2023-05-01T00:00:00Z") // 1682899200
	suite.ctx = suite.ctx.WithBlockTime(t)
	suite.ctx = suite.ctx.WithBlockHeight(prevoteHeight)

	prevoteBlockHight := suite.ctx.BlockHeight()

	//PREVOTE
	suite.Prevote(suite.addrs[0], "abcdef", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5,1:1682899198:BTC:28110.45,4:1682899198:BTC:28108.25")
	suite.Prevote(suite.addrs[1], "123456", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5,2:1682899197:BTC:28105.92,5:1682899198:BTC:28107.55")
	suite.Prevote(suite.addrs[2], "ababab", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5,3:1682899196:BTC:28108.11,6:1682899194:BTC:28108.65")
	suite.Prevote(suite.addrs[3], "dedede", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5,1:1682899198:BTC:28110.43,4:1682899198:BTC:28108.75")
	suite.Prevote(suite.addrs[4], "aaaaaa", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5,2:1682899197:BTC:28105.94,5:1682899198:BTC:28107.85")
	suite.Prevote(suite.addrs[5], "bbbbbb", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5,3:1682899196:BTC:28108.15,6:1682899194:BTC:28108.95")

	//oracle.EndBlocker(suite.ctx, suite.keeper)

	//VOTE BLOCK SET
	t2 := types.ParseTime("2023-05-01T00:00:05Z") // 1682899205	// +5
	suite.ctx = suite.ctx.WithBlockTime(t2)
	suite.ctx = suite.ctx.WithBlockHeight(prevoteHeight + 1)

	//VOTE
	suite.Vote(suite.addrs[0], "abcdef", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5,1:1682899198:BTC:28110.45,4:1682899198:BTC:28108.25", prevoteBlockHight)
	suite.Vote(suite.addrs[1], "123456", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5,2:1682899197:BTC:28105.92,5:1682899198:BTC:28107.55", prevoteBlockHight)
	suite.Vote(suite.addrs[2], "ababab", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5,3:1682899196:BTC:28108.11,6:1682899194:BTC:28108.65", prevoteBlockHight)
	suite.Vote(suite.addrs[3], "dedede", "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5,1:1682899198:BTC:28110.43,4:1682899198:BTC:28108.75", prevoteBlockHight)
	suite.Vote(suite.addrs[4], "aaaaaa", "2:1682899197:ETH:1805.9,5:1682899198:ETH:1807.5,2:1682899197:BTC:28105.94,5:1682899198:BTC:28107.85", prevoteBlockHight)
	suite.Vote(suite.addrs[5], "bbbbbb", "3:1682899196:ETH:1808.1,6:1682899194:ETH:1808.5,3:1682899196:BTC:28108.15,6:1682899194:BTC:28108.95", prevoteBlockHight)

	oracle.MidBlocker(suite.ctx, suite.keeper)

	//TODO: delete. debug
	price := suite.querier.Keeper.AllOraclePrices(suite.ctx)
	fmt.Println(price)

	/*
		[ETH]
		median 1808.300000000000000000
		MAD 0.500000000000000000
		low 1807.050000000000000000
		high 1809.550000000000000000
		price 1808.150000000000000000

		[BTC]
		median 28108.200000000000000000
		MAD 0.600000000000000000
		low 28106.700000000000000000
		high 28109.700000000000000000
		price 28108.282500000000000000
	*/
	suite.Require().Equal(2, len(price))
	suite.Require().Equal("BTC", price[0].Ticker)
	suite.Require().Equal("28108.282500000000000000", price[0].UsdPrice.String()) // 18 precision
	suite.Require().Equal("ETH", price[1].Ticker)
	suite.Require().Equal("1808.150000000000000000", price[1].UsdPrice.String()) // 18 precision

}
