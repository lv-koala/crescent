package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

// IterateAllMedianPrices iterates over all median prices.
// Iterator stops when exhausting the source, or when the handler returns `true`.
func (k Keeper) IterateAllOraclePrices(
	ctx sdk.Context,
	handler func(types.OraclePrice) bool,
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PriceKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var op types.OraclePrice
		//var decProto sdk.DecProto
		k.cdc.MustUnmarshal(iter.Value(), &op)
		//denom, blockNum := types.ParseDenomAndBlockFromKey(iter.Key(), types.KeyPrefixMedian)

		if handler(op) {
			break
		}
	}
}

// AllMedianPrices is a helper function that collects and returns all
// median prices using the IterateAllMedianPrices iterator
func (k Keeper) AllOraclePrices(ctx sdk.Context) types.OraclePrices {
	prices := types.OraclePrices{}
	k.IterateAllOraclePrices(ctx, func(op types.OraclePrice) (stop bool) {
		prices = append(prices, op)
		return false
	})
	return prices
}
