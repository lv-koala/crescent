package types

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Oracle Errors
var (
	ErrInvalidPrices           = sdkerrors.Register(ModuleName, 2, "invalid exchange rate")
	ErrNoPrevote               = sdkerrors.Register(ModuleName, 3, "no prevote")
	ErrNoVote                  = sdkerrors.Register(ModuleName, 4, "no vote")
	ErrNoVotingPermission      = sdkerrors.Register(ModuleName, 5, "unauthorized voter")
	ErrInvalidHash             = sdkerrors.Register(ModuleName, 6, "invalid hash")
	ErrInvalidHashLength       = sdkerrors.Register(ModuleName, 7, fmt.Sprintf("invalid hash length; should equal %d", tmhash.TruncatedSize))
	ErrVerificationFailed      = sdkerrors.Register(ModuleName, 8, "hash verification failed")
	ErrRevealPeriodMissMatch   = sdkerrors.Register(ModuleName, 9, "reveal period of submitted vote do not match with registered prevote")
	ErrInvalidSaltLength       = sdkerrors.Register(ModuleName, 10, "invalid salt length; should be 1~4")
	ErrNoAggregatePrevote      = sdkerrors.Register(ModuleName, 11, "no aggregate prevote")
	ErrNoAggregateVote         = sdkerrors.Register(ModuleName, 12, "no aggregate vote")
	ErrNoTobinTax              = sdkerrors.Register(ModuleName, 13, "no tobin tax")
	ErrUnknownTicker           = sdkerrors.Register(ModuleName, 14, "unknown ticker")
	ErrInvalidOraclePrice      = sdkerrors.Register(ModuleName, 15, "invalid oracle price")
	ErrNotAllowedFeeder        = sdkerrors.Register(ModuleName, 16, "no permission to feed")
	ErrExistingPrevote         = sdkerrors.Register(ModuleName, 17, "prevote already submitted for same height")
	ErrAlreadyRegisteredTicker = sdkerrors.Register(ModuleName, 18, "ticker already registered")
	ErrInvalidTicker           = sdkerrors.Register(ModuleName, 19, "invalid ticker")
	ErrUpdateFeeder            = sdkerrors.Register(ModuleName, 20, "invalid feeder update")
)
