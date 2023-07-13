package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// AccountKeeper is expected keeper for auth module
type AccountKeeper interface {
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
	GetModuleAddress(name string) sdk.AccAddress
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI // only used for simulation
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)

	// only used for simulation
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
