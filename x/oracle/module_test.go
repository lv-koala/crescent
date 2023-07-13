package oracle_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "github.com/crescent-network/crescent/v5/app"
	"github.com/crescent-network/crescent/v5/cmd/crescentd/cmd"
	"github.com/crescent-network/crescent/v5/x/oracle/keeper"
	types "github.com/crescent-network/crescent/v5/x/oracle/types"
)

var (
	initialBalances = sdk.NewCoins(
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 1_000_000_000))
)

type ModuleTestSuite struct {
	suite.Suite

	app     *chain.App
	ctx     sdk.Context
	keeper  keeper.Keeper
	querier keeper.Querier
	addrs   []sdk.AccAddress
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	cmd.GetConfig()
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	suite.app = app
	suite.ctx = ctx
	suite.keeper = suite.app.OracleKeeper
	suite.querier = keeper.Querier{Keeper: suite.keeper}
	suite.addrs = chain.AddTestAddrs(suite.app, suite.ctx, 6, sdk.ZeroInt())
	for _, addr := range suite.addrs {
		err := chain.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
	params := suite.keeper.GetParams(suite.ctx)
	params.AcceptTickers = []string{"ETH", "BTC", "ATOM"}
	params.AcceptableSeconds = 60

	params.WhitelistFeeders = []string{suite.addrs[0].String(), suite.addrs[1].String(), suite.addrs[2].String(), suite.addrs[3].String(), suite.addrs[4].String(), suite.addrs[5].String()}
	suite.keeper.SetParams(suite.ctx, params)

}

// Prevote is a convenient method to test Keeper.Prevote.
func (suite *ModuleTestSuite) Prevote(feederAcc sdk.AccAddress, salt string, voteString string) {

	var err error
	if !types.ValidatePriceTuplesString(voteString) {
		err = fmt.Errorf("invalid exchange rates string")
	}
	suite.Require().NoError(err)

	hash := types.GetAggregateVoteHash(salt, voteString, feederAcc)

	msg := types.NewMsgPricesPrevote(hash.String(), feederAcc)
	suite.Require().NoError(msg.ValidateBasic())

	err = suite.keeper.HandlePrevoteMsg(suite.ctx, msg)
	suite.Require().NoError(err)
}

// Rewards is a convenient method to test Keeper.WithdrawAllRewards.
func (suite *ModuleTestSuite) Vote(feederAcc sdk.AccAddress, salt string, voteString string, height int64) {
	//cacheCtx, _ := suite.ctx.CacheContext()

	var err error
	if !types.ValidatePriceTuplesString(voteString) {
		err = fmt.Errorf("invalid exchange rates string")
	}
	suite.Require().NoError(err)

	msg := types.NewMsgPricesVote(salt, voteString, feederAcc)
	suite.Require().NoError(msg.ValidateBasic())

	err = suite.keeper.HandleVoteMsg(suite.ctx, msg)
	suite.Require().NoError(err)
}
