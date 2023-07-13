package oracle

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/v5/x/oracle/keeper"
	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

func MidBlocker(ctx sdk.Context, k keeper.Keeper) {

	// vote msg will be processed before midblock
	// calc price using valid votes
	params := k.GetParams(ctx)
	CalcPrices(ctx, params, k)

	//Delete all prevotes and validVotes

}

func CalcPrices(ctx sdk.Context, params types.Params, k keeper.Keeper) {

	// missing check
	mapFeeder := make(map[string]bool)
	for _, v := range params.WhitelistFeeders {
		mapFeeder[v] = false
	}

	// voteTargets defines the symbol (ticker) denoms that we require votes on
	for _, v := range params.AcceptTickers {
		validVotes, err := k.GetValidVotes(ctx, v)
		if err != nil {
			//skip empty validVotes
			continue
		}

		median, err := validVotes.ValidTuple.Median()
		if err != nil {
			//skip empty validVotes
			continue
		}

		//TODO check out the median of the validVotes
		m, err := validVotes.ValidTuple.MAD()
		if err != nil {
			//skip empty validVotes
			continue
		}

		mad25 := m.Mul(sdk.NewDecWithPrec(25, 1))

		sum := sdk.ZeroDec()
		ll := 0
		//check missing feeder
		LowBound := median.Sub(mad25)
		HighBound := median.Add(mad25)

		for _, v := range validVotes.ValidTuple {
			if v.UsdPrice.LTE(HighBound) && v.UsdPrice.GTE(LowBound) {
				sum = sum.Add(v.UsdPrice)
				ll++
				mapFeeder[v.Feeder] = true
			}
		}

		sum = sum.QuoInt64(int64(ll))

		//TODO: delete
		fmt.Println("median", median.String())
		fmt.Println("MAD", m.String())
		fmt.Println("low", LowBound.String())
		fmt.Println("high", HighBound.String())
		fmt.Println("price", sum.String())

		//set price
		//The price is from the previous block(first from prevote), but it's currently available, so it's recorded as the current block's price
		k.SetPrice(ctx, v, ctx.BlockHeight(), &sum)

		//clear validvotes
		//k.ClearVotes(ctx, params.VotePeriod)

	}

	for f, v := range mapFeeder {
		if !v {
			addr := sdk.AccAddress(f)
			old := k.GetMissCounter(ctx, addr)
			//TODO add miss counter
			k.SetMissCounter(ctx, addr, old+1)
		}
	}

}
