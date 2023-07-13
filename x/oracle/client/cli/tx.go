package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

// GetTxCmd returns the CLI transaction commands for the x/oracle module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transaction commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdPrevote(),
		GetCmdVote(),
	)

	return cmd
}

// GetCmdAggregatePricesPrevote creates a Cobra command to generate or
// broadcast a transaction with a MsgPricePrevote message.
func GetCmdPrevote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prevote [salt-hex-encoded] [price-tuples] [feeder]",
		Args:  cobra.RangeArgs(2, 3),
		Short: "Submit an prevote with a hash",
		Long: fmt.Sprintf(`Submit a price prevote with a hash as a hex string
			representation of a byte array.
			Ex: crescentd tx oracle prevote %s %s %s`,
			"19c30cf9ea8aa0e0b03904162cadec0f2024a76d",
			"1:14612312311:ETH:1800.3,3:1461231313131:ETH:1802.3",
			"cre1zaavvzxez0elundtn32qnk9lkm8kmcszxclz6p",
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			salt := args[0]
			PricessStr := args[1]

			feederAddress := sdk.AccAddress(clientCtx.GetFromAddress())
			if len(args) > 2 {
				feederAddress, err = sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return err
				}
			}

			if !types.ValidatePriceTuplesString(PricessStr) {
				return fmt.Errorf("invalid exchange rates string")
			}

			hash := types.GetAggregateVoteHash(salt, PricessStr, feederAddress)

			//clientCtx.GetFromAddress(),
			msg := types.NewMsgPricesPrevote(
				hash.String(),
				feederAddress,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAggregatePricesVote creates a Cobra command to generate or
// broadcast a transaction with a NewMsgAggregatePricesVote message.
func GetCmdVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote [salt] [exchange-rates] [feeder]",
		Args:  cobra.RangeArgs(2, 3),
		Short: "Submit an exchange rate vote with the salt and exchange rate string",
		Long: fmt.Sprintf(`Submit an exchange rate vote with the salt of the previous hash, and the
			exchange rate string previously used in the hash.
			Ex: crescentd tx oracle vote %s %s --from alice`,
			"0cf33fb528b388660c3a42c3f3250e983395290b75fef255050fb5bc48a6025f",
			"1:14612312311:ETH:1800.3,3:1461231313131:ETH:1802.3",
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			salt := args[0]
			PricessStr := args[1]

			feederAddress := sdk.AccAddress(clientCtx.GetFromAddress())
			if len(args) > 2 {
				feederAddress, err = sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return err
				}
			}

			if !types.ValidatePriceTuplesString(PricessStr) {
				return fmt.Errorf("invalid exchange rates string")
			}

			//clientCtx.GetFromAddress(),
			msg := types.NewMsgPricesVote(
				salt,
				PricessStr,
				feederAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
