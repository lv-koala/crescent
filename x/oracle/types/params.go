package types

import (
	"fmt"
	"strings"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyVoteThreshold            = []byte("VoteThreshold")
	KeyRewardRate               = []byte("RewardRate")
	KeyRewardDistributionPeriod = []byte("RewardDistributionPeriod")
	KeyWhitelistFeeders         = []byte("WhitelistFeeders")
	KeyDelegateAccounts         = []byte("DelegateAccounts")
	KeyAcceptTickers            = []byte("AcceptTickers")
	KeyMaxAcceptLatencySeconds  = []byte("MaxAcceptLatencySeconds")
)

var _ paramstypes.ParamSet = &Params{}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params {

	return Params{
		VoteMinTx:                3,
		RewardDistributionRate:   10,  //10%
		RewardDistributionPeriod: 100, //100 blocks
		WhitelistFeeders:         []string{},
		DelegateAccounts:         []string{},
		AcceptTickers:            []string{},
		AcceptableSeconds:        10, // 10 seconds
	}
}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of oracle module's parameters.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(
			KeyVoteThreshold,
			&p.VoteMinTx,
			validateVoteThreshold,
		),
		paramstypes.NewParamSetPair(
			KeyRewardRate,
			&p.RewardDistributionRate,
			validateRewardDistributionRate,
		),
		paramstypes.NewParamSetPair(
			KeyRewardDistributionPeriod,
			&p.RewardDistributionPeriod,
			validateRewardDistributionPeriod,
		),
		paramstypes.NewParamSetPair(
			KeyWhitelistFeeders,
			&p.WhitelistFeeders,
			validateWhitelistFeeders,
		),
		paramstypes.NewParamSetPair(
			KeyDelegateAccounts,
			&p.DelegateAccounts,
			validateDelegateAccounts,
		),
		paramstypes.NewParamSetPair(
			KeyAcceptTickers,
			&p.AcceptTickers,
			validateAcceptSymbols,
		),
		paramstypes.NewParamSetPair(
			KeyMaxAcceptLatencySeconds,
			&p.AcceptableSeconds,
			validateMaxAcceptLatencySeconds,
		),
	}
}

// Validate performs basic validation on oracle parameters.
func (p Params) Validate() error {
	if p.VoteMinTx == 0 {
		return fmt.Errorf("oracle parameter VoteMin must be > 0, is %d", p.VoteMinTx)
	}
	if p.RewardDistributionRate > 100 {
		return fmt.Errorf("oracle parameter RewardDistributionRate must be less than 100 percent")
	}

	if p.AcceptableSeconds < 2 {
		return fmt.Errorf("oracle vote should be larger than 2")
	}

	for _, s := range p.AcceptTickers {
		if len(s) == 0 {
			return fmt.Errorf("oracle accept symbol must not empty")
		}
	}

	for _, s := range p.WhitelistFeeders {
		if len(s) == 0 {
			return fmt.Errorf("feeder must not empty")
		}
	}

	return nil
}

func validateVoteThreshold(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("vote period must be positive: %d", v)
	}

	return nil
}

func validateRewardDistributionRate(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("vote period must be positive: %d", v)
	}

	if v > 100 {
		return fmt.Errorf("vote period must be less than 100: %d", v)
	}

	return nil
}

func validateRewardDistributionPeriod(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateWhitelistFeeders(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, d := range v {
		if len(d) == 0 {
			return fmt.Errorf("feeder address has empty")
		}
		/*if !strings.HasPrefix(d, "cre") {
			return fmt.Errorf("invalid feeder address: %s", d)
		}
		*/
	}

	return nil
}

func validateDelegateAccounts(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, d := range v {
		if len(d) == 0 {
			return fmt.Errorf("delegation address has empty")
		}
		/*if !strings.HasPrefix(d, "cre") {
			return fmt.Errorf("invalid feeder address: %s", d)
		}
		*/
	}

	return nil
}

func validateAcceptSymbols(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, d := range v {
		if len(d) == 0 {
			return fmt.Errorf("oracle parameter AcceptSymbol has empty ")
		}
		if strings.ToUpper(d) != d {
			return fmt.Errorf("oracle parameter AcceptSymbol must be upper case: %s", d)
		}
	}

	return nil
}

func validateMaxAcceptLatencySeconds(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("accept time must be positive: %d", v)
	}

	if v < 2 {
		return fmt.Errorf("vote period must be larger than 2: %d", v)
	}

	return nil
}
