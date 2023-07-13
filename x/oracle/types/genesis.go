package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params, missCounters []MissCounter,
	pricePrevotes []PricePrevote,
	priceVotes []PriceVote,
	priceOracle OraclePrices,
) *GenesisState {
	return &GenesisState{
		Params:        params,
		MissCounters:  missCounters,
		PricePrevotes: pricePrevotes,
		PriceVotes:    priceVotes,
		PriceOracle:   priceOracle,
	}
}

// DefaultGenesisState - default GenesisState used by columbus-2
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams(),
		[]MissCounter{},
		[]PricePrevote{},
		[]PriceVote{},
		[]OraclePrice{})
}

// ValidateGenesis validates the oracle genesis state
func ValidateGenesis(data *GenesisState) error {
	return data.Params.Validate()
}

// GetGenesisStateFromAppState returns x/oracle GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
