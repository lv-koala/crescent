package keeper

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

// Keeper of the oracle store
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	tsKey      sdk.StoreKey
	paramSpace paramstypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper

	//distrName string

}

// NewKeeper constructs a new keeper for oracle
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	tsKey sdk.StoreKey,
	paramspace paramstypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	// ensure oracle module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// set KeyTable if it has not already been set
	if !paramspace.HasKeyTable() {
		paramspace = paramspace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		tsKey:         tsKey,
		paramSpace:    paramspace,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetTsKey() sdk.StoreKey {
	return k.tsKey
}

// GetPrice gets the latest consensus price USD denominated in the
// denom asset from the store.
func (k Keeper) GetPrice(ctx sdk.Context, symbol string) (types.OraclePrice, error) {
	store := ctx.KVStore(k.storeKey)
	symbol = strings.ToUpper(symbol)
	b := store.Get(types.GetPriceKey(symbol))
	p := types.OraclePrice{}
	if b == nil {
		return p, types.ErrUnknownTicker.Wrap(symbol)
	}

	k.cdc.MustUnmarshal(b, &p)

	return p, nil
}

// SetPrices sets the consensus exchange rate of USD denominated in the
// denom asset to the store.
func (k Keeper) SetPrice(ctx sdk.Context, ticker string, prevoteHeight int64, price *sdk.Dec) {
	store := ctx.KVStore(k.storeKey)

	//check symbol is valid

	bz := k.cdc.MustMarshal(&types.OraclePrice{Ticker: ticker, UsdPrice: *price, BlockNum: prevoteHeight})
	ticker = strings.ToUpper(ticker)
	store.Set(types.GetPriceKey(ticker), bz)
}

// IterateMissCounters iterates over the miss counters and performs a callback
// function.
func (k Keeper) IteratePrices(ctx sdk.Context, handler func(types.OraclePrice) bool) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.PriceKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		//ticker := iter.Key()[2:]

		var pr types.OraclePrice
		k.cdc.MustUnmarshal(iter.Value(), &pr)

		if handler(pr) {
			break
		}
	}
}

// GetMissCounter retrieves the # of vote periods missed in this oracle slash
// window.
func (k Keeper) GetMissCounter(ctx sdk.Context, operator sdk.AccAddress) uint64 {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetMissCounterKey(operator))
	if bz == nil {
		// by default the counter is zero
		return 0
	}

	var missCounter gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &missCounter)

	return missCounter.Value
}

// SetMissCounter updates the # of vote periods missed in this oracle slash
// window.
func (k Keeper) SetMissCounter(ctx sdk.Context, operator sdk.AccAddress, missCounter uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: missCounter})
	store.Set(types.GetMissCounterKey(operator), bz)
}

// DeleteMissCounter removes miss counter for the validator.
func (k Keeper) DeleteMissCounter(ctx sdk.Context, operator sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMissCounterKey(operator))
}

// IterateMissCounters iterates over the miss counters and performs a callback
// function.
func (k Keeper) IterateMissCounters(ctx sdk.Context, handler func(sdk.AccAddress, uint64) bool) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.MissCounterKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		operator := sdk.AccAddress(iter.Key()[2:])

		var missCounter gogotypes.UInt64Value
		k.cdc.MustUnmarshal(iter.Value(), &missCounter)

		if handler(operator, missCounter.Value) {
			break
		}
	}
}

func (k Keeper) SetLsvFeeder(ctx sdk.Context, lsvAccAddr string, feeder string) {
	store := ctx.KVStore(k.storeKey)

	f := types.LsvFeeder{LsvAddr: lsvAccAddr, FeederAddr: feeder, UpdateHeight: ctx.BlockHeight()}
	bz := k.cdc.MustMarshal(&f)
	store.Set(types.GetLsvFeederKey(lsvAccAddr), bz)
}

// GetAggregatePricesPrevote retrieves an oracle prevote from the store.
func (k Keeper) GetPrevoteWithHeight(
	ctx sdk.Context,
	feeder sdk.AccAddress,
	height uint64,
) (types.PricePrevote, error) {
	store := ctx.KVStore(k.storeKey)

	key := types.GetPrevoteKeyWithHeight(feeder, height)
	bz := store.Get(key)
	if bz == nil {
		return types.PricePrevote{}, fmt.Errorf("prevote not found for feeder %s with height %d", feeder, height)
	}

	var aggregatePrevote types.PricePrevote
	k.cdc.MustUnmarshal(bz, &aggregatePrevote)

	return aggregatePrevote, nil
}

// HasAggregatePricesPrevote checks if a validator has an existing prevote.
func (k Keeper) HasAggregatePricesPrevoteWithHeight(
	ctx sdk.Context,
	feeder sdk.AccAddress,
	height uint64,
) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetPrevoteKeyWithHeight(feeder, height))
}

// SetAggregatePricesPrevote set an oracle aggregate prevote to the store.
func (k Keeper) SetAggregatePricesPrevote(
	ctx sdk.Context,
	feeder sdk.AccAddress,
	prevote types.PricePrevote,
) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&prevote)
	store.Set(types.GetPrevoteKeyWithHeight(feeder, uint64(prevote.Height)), bz)
}

// DeleteAggregatePricesPrevote deletes an oracle prevote from the store.
func (k Keeper) DeleteAggregatePricesPrevote(ctx sdk.Context, feeder sdk.AccAddress, height uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPrevoteKeyWithHeight(feeder, height))
}

// IterateAggregatePricesPrevotes iterates rate over prevotes in the store
func (k Keeper) IterateAggregatePricesPrevotes(
	ctx sdk.Context,
	handler func(sdk.AccAddress, types.PricePrevote) bool,
) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.PrevoteKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		//TODO: addr length
		feederAddr := sdk.AccAddress(iter.Key()[2:42])

		var aggregatePrevote types.PricePrevote
		k.cdc.MustUnmarshal(iter.Value(), &aggregatePrevote)

		if handler(feederAddr, aggregatePrevote) {
			break
		}
	}
}

// VOTE is not used. If vote msg
// GetAggregatePricesPrevote retrieves an oracle prevote from the store.
// valid votes are stored in transient store
func (k Keeper) GetValidVotes(
	ctx sdk.Context,
	symbol string,
) (types.ValidVotes, error) {
	//store := ctx.KVStore(k.storeKey)
	tstore := ctx.TransientStore(k.tsKey)

	bz := tstore.Get(types.GetValidVoteKey(symbol))
	if bz == nil {
		return types.ValidVotes{}, fmt.Errorf("not found for symbol %s", symbol)
	}

	var v types.ValidVotes
	k.cdc.MustUnmarshal(bz, &v)

	return v, nil
}

// SetAggregatePricesPrevote set an oracle aggregate prevote to the store.
func (k Keeper) SetValidVotes(
	ctx sdk.Context,
	symbol string,
	validVote types.ValidVotes,
) {
	//store := ctx.KVStore(k.storeKey)
	tstore := ctx.TransientStore(k.tsKey)

	bz := k.cdc.MustMarshal(&validVote)
	tstore.Set(types.GetValidVoteKey(symbol), bz)
}

// ValidateFeeder returns error if the given feeder is not allowed to feed the message.
func (k Keeper) ValidateFeeder(ctx sdk.Context, feederAddr sdk.AccAddress) error {
	params := k.GetParams(ctx)

	checkFeederAddr := feederAddr.String()
	for _, feeder := range params.WhitelistFeeders {
		if feeder == checkFeederAddr {
			return nil
		}
	}

	return types.ErrNotAllowedFeeder.Wrap(checkFeederAddr)

	//return nil
}

// ValidateFeeder returns error if the given account is not in delegateAccounts
func (k Keeper) ValidateDelegateAccount(ctx sdk.Context, delegateAddr sdk.AccAddress) error {
	params := k.GetParams(ctx)

	checkFeederAddr := delegateAddr.String()
	for _, feeder := range params.DelegateAccounts {
		if feeder == checkFeederAddr {
			return nil
		}
	}

	return types.ErrNotAllowedFeeder.Wrap(checkFeederAddr)

	//return nil
}

// ValidateFeeder returns error if the given feeder is not allowed to feed the message.
func (k Keeper) ValidateTicker(ctx sdk.Context, inputTicker string) error {
	params := k.GetParams(ctx)

	for _, symbol := range params.AcceptTickers {
		if symbol == inputTicker {
			return nil
		}
	}

	return types.ErrUnknownTicker.Wrap(inputTicker)

	//return nil
}

// CreatePair handles types.MsgCreatePair and creates a pair.
func (k Keeper) HandlePrevoteMsg(ctx sdk.Context, msg *types.MsgPricesPrevote) error {

	fmt.Println("PREVOTE MSG", msg)
	feederAcc, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		return err
	}

	errFeeder := k.ValidateFeeder(ctx, feederAcc)
	errDelegateAddr := k.ValidateDelegateAccount(ctx, feederAcc)
	if errFeeder != nil && errDelegateAddr != nil {
		return errFeeder
	}

	// Ensure prevote wasn't already submitted
	if k.HasAggregatePricesPrevoteWithHeight(ctx, feederAcc, uint64(ctx.BlockHeight())) {
		return types.ErrExistingPrevote
	}

	// Convert hex string to votehash
	voteHash, err := types.AggregateVoteHashFromHex(msg.Hash)
	if err != nil {
		return types.ErrInvalidHash.Wrap(err.Error())
	}

	aggregatePrevote := types.NewPricePrevote(voteHash, feederAcc, ctx.BlockHeight())
	k.SetAggregatePricesPrevote(ctx, feederAcc, aggregatePrevote)

	//debug
	fmt.Println("prevote", aggregatePrevote)

	//TODO: emit event
	/*
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypePricesPrevote,
				sdk.NewAttribute(types.AttributeKeyCreator, msg.feeder),
				sdk.NewAttribute(types.AttributeKeyBaseCoinDenom, msg.BaseCoinDenom),
				sdk.NewAttribute(types.AttributeKeyQuoteCoinDenom, msg.QuoteCoinDenom),
				sdk.NewAttribute(types.AttributeKeyPairId, strconv.FormatUint(pair.Id, 10)),
				sdk.NewAttribute(types.AttributeKeyEscrowAddress, pair.EscrowAddress),
			),
		})
	*/

	return nil

}

func (k Keeper) HandleVoteMsg(ctx sdk.Context, msg *types.MsgPricesVote) error {

	fmt.Println("VOTE MSG", msg)
	feederAddr, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		fmt.Println("invalid addr: ", msg.Feeder)
		return err
	}

	errFeeder := k.ValidateFeeder(ctx, feederAddr)
	errDelegateAddr := k.ValidateDelegateAccount(ctx, feederAddr)
	if errFeeder != nil && errDelegateAddr != nil {
		fmt.Println("invalid feeder: ", msg.Feeder)
		return errFeeder
	}

	prevBlock := uint64(ctx.BlockHeight()) - 1
	if prevBlock < 1 {
		fmt.Println("prev: ", prevBlock)
		return types.ErrNoPrevote
	}
	aggregatePrevote, err := k.GetPrevoteWithHeight(ctx, feederAddr, prevBlock)
	if err != nil {
		fmt.Println("no pre", prevBlock)
		return types.ErrNoAggregatePrevote.Wrap(msg.Feeder)
	}

	priceTuples, err := types.ParsePriceTuples(msg.PriceTuples)
	if err != nil {
		fmt.Println("invalid Pricess: ", msg.PriceTuples)
		return err
	}

	fmt.Println("tuples: ", len(priceTuples))
	fmt.Println(msg.Salt, msg.PriceTuples, feederAddr)

	// Verify that the vote hash and prevote hash match
	// salt should be hex string
	hash := types.GetAggregateVoteHash(msg.Salt, msg.PriceTuples, feederAddr)
	if aggregatePrevote.Hash != hash.String() {
		fmt.Println("hash err: ", aggregatePrevote.Hash, hash)
		return types.ErrVerificationFailed.Wrapf("must be given %s not %s", aggregatePrevote.Hash, hash)
	}

	params := k.GetParams(ctx)

	validSymbolCount := 0
	for _, tuple := range priceTuples {
		for _, s := range params.AcceptTickers {
			if tuple.Ticker == s {
				validSymbolCount++
				break
			}
		}
	}

	fmt.Println("valid vote tuple: ", validSymbolCount, len(priceTuples))

	// if any unknown symbol included, return error
	if validSymbolCount < len(priceTuples) {
		return types.ErrUnknownTicker
	}

	// Filter out rates which aren't included in the AcceptList
	// If they missed some, then Increases the miss counter, which affects reward distribution.

	// Only timeout tuples are removed.
	for _, tuple := range priceTuples {
		timeBefore := tuple.Timestamp + params.AcceptableSeconds
		t := time.Unix(int64(timeBefore), 0)
		if !ctx.BlockTime().After(t) { // BlockTime <= collectTime + acceptableSeconds
			v := types.ValidPriceTuple{
				SourceExchangeId: tuple.SourceExchangeId,
				Feeder:           feederAddr.String(),
				UsdPrice:         tuple.UsdPrice,
			}
			valueOld, err := k.GetValidVotes(ctx, tuple.Ticker)
			if err != nil {
				// new symbol.
				// height is not prevoted height, but voted height
				valueOld.Height = ctx.BlockHeight()
			} else {
				if valueOld.Height != ctx.BlockHeight() {
					return types.ErrVerificationFailed.Wrapf("validVote always has same height. validvote Height %d not blockHeight %d", valueOld.Height, ctx.BlockHeight())
				}
			}
			valueOld.ValidTuple = append(valueOld.ValidTuple, v)
			valueOld.ValidTupleCount = int64(len(valueOld.ValidTuple))
			k.SetValidVotes(ctx, tuple.Ticker, valueOld) //overwrite
		}
	}

	return nil
}

func (k Keeper) HandleAddNewTickerMsg(ctx sdk.Context, msg *types.MsgAddNewTicker) error {

	fmt.Println("AddNewTicker MSG", msg)
	feederAddr, err := sdk.AccAddressFromBech32(msg.DelegateAddress)
	if err != nil {
		fmt.Println("invalid addr: ", msg.DelegateAddress)
		return err
	}

	if len(msg.Ticker) < 3 {
		return types.ErrInvalidTicker
	}

	errDelegateAddr := k.ValidateDelegateAccount(ctx, feederAddr)
	if errDelegateAddr != nil {
		return errDelegateAddr
	}

	params := k.GetParams(ctx)

	t := strings.ToUpper(msg.Ticker)
	// Check if the ticker is already registered
	for _, s := range params.AcceptTickers {
		if s == t {
			return types.ErrAlreadyRegisteredTicker
		}
	}

	params.AcceptTickers = append(params.AcceptTickers, t)
	k.SetParams(ctx, params)

	fmt.Println("acceptTickers: ", params.AcceptTickers)

	return nil
}

func (k Keeper) HandleLsvFeederUpdateMsg(ctx sdk.Context, msg *types.MsgLsvFeederUpdate) error {

	fmt.Println("Lsv Update", msg)
	feederAddr, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		fmt.Println("invalid addr: ", msg.Feeder)
		return err
	}

	//check
	params := k.GetParams(ctx)

	// Check if the ticker is already registered
	for _, s := range params.WhitelistFeeders {
		if s == string(feederAddr) {
			return types.ErrUpdateFeeder
		}
	}

	oldFeeder := types.GetLsvFeederKey(msg.LsvAccAddr)

	//TODO: check LSV address
	if len(oldFeeder) > 0 { // replace if exists
		for i, s := range params.WhitelistFeeders {
			if s == string(oldFeeder) {
				params.WhitelistFeeders[i] = string(feederAddr)
				break
			}
		}
	} else {
		params.WhitelistFeeders = append(params.WhitelistFeeders, string(feederAddr))
	}
	k.SetWhitelistFeeders(ctx, params.WhitelistFeeders)

	k.SetLsvFeeder(ctx, msg.LsvAccAddr, msg.Feeder)

	fmt.Println("lsv updated")

	return nil
}
