package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramstypes.ParamSet = (*Params)(nil)

var (
	KeyPoolCreationFee               = []byte("PoolCreationFee")
	KeyDefaultTickSpacing            = []byte("DefaultTickSpacing")
	KeyPrivateFarmingPlanCreationFee = []byte("PrivateFarmingPlanCreationFee")
	KeyMaxNumPrivateFarmingPlans     = []byte("MaxNumPrivateFarmingPlans")
	KeyMaxRewardsBlockTime           = []byte("MaxRewardsBlockTime")
)

var (
	DefaultPoolCreationFee               = sdk.NewCoins()
	DefaultDefaultTickSpacing            = uint32(50)
	DefaultPrivateFarmingPlanCreationFee = sdk.NewCoins()
	DefaultMaxNumPrivateFarmingPlans     = uint32(50)
	DefaultMaxRewardsBlockTime           = 10 * time.Second
)

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default params for the module.
func DefaultParams() Params {
	return Params{
		PoolCreationFee:               DefaultPoolCreationFee,
		DefaultTickSpacing:            DefaultDefaultTickSpacing,
		PrivateFarmingPlanCreationFee: DefaultPrivateFarmingPlanCreationFee,
		MaxNumPrivateFarmingPlans:     DefaultMaxNumPrivateFarmingPlans,
		MaxRewardsBlockTime:           DefaultMaxRewardsBlockTime,
	}
}

// ParamSetPairs implements ParamSet.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyPoolCreationFee, &params.PoolCreationFee, validatePoolCreationFee),
		paramstypes.NewParamSetPair(KeyDefaultTickSpacing, &params.DefaultTickSpacing, validateDefaultTickSpacing),
		paramstypes.NewParamSetPair(KeyPrivateFarmingPlanCreationFee, &params.PrivateFarmingPlanCreationFee, validatePrivateFarmingPlanCreationFee),
		paramstypes.NewParamSetPair(KeyMaxNumPrivateFarmingPlans, &params.MaxNumPrivateFarmingPlans, validateMaxNumPrivateFarmingPlans),
		paramstypes.NewParamSetPair(KeyMaxRewardsBlockTime, &params.MaxRewardsBlockTime, validateMaxRewardsBlockTime),
	}
}

// Validate validates Params.
func (params Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{params.PoolCreationFee, validatePoolCreationFee},
		{params.DefaultTickSpacing, validateDefaultTickSpacing},
		{params.PrivateFarmingPlanCreationFee, validatePrivateFarmingPlanCreationFee},
		{params.MaxNumPrivateFarmingPlans, validateMaxNumPrivateFarmingPlans},
		{params.MaxRewardsBlockTime, validateMaxRewardsBlockTime},
	} {
		if err := field.validateFunc(field.val); err != nil {
			return err
		}
	}
	return nil
}

func validatePoolCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid pool creation fee: %w", err)
	}
	return nil
}

func validateDefaultTickSpacing(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v == 0 {
		return fmt.Errorf("tick spacing must be positive")
	}
	return nil
}

func validatePrivateFarmingPlanCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid private farming plan creation fee: %w", err)
	}
	return nil
}

func validateMaxNumPrivateFarmingPlans(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateMaxRewardsBlockTime(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("max rewards block time must be positive")
	}
	return nil
}
