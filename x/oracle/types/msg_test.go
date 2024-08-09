package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	utils "github.com/crescent-network/crescent/v5/types"
	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

func TestMsgMsgPrevote_ValidateBasic(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(msg *types.MsgPricesPrevote)
		expectedErr string
	}{
		{
			"valid",
			func(msg *types.MsgPricesPrevote) {},
			"",
		},
		{
			"invalid feeder",
			func(msg *types.MsgPricesPrevote) {
				msg.Feeder = "invalidaddr"
			},
			"invalid sender address: decoding bech32 failed: invalid separator index -1: invalid address",
		},
		{
			"invalid feeder",
			func(msg *types.MsgPricesPrevote) {
				msg.Feeder = "cre1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8rhheuj"
			},
			"invalid sender address: decoding bech32 failed: invalid separator index -1: invalid address",
		},
		{
			"invalid hash",
			func(msg *types.MsgPricesPrevote) {
				msg.Hash = ""
			},
			"invalid hash length; should equal 20",
		},
		{
			"invalid hash",
			func(msg *types.MsgPricesPrevote) {
				msg.Hash = "qwertyuiopasdfghjklzxcvbnm"
				// invalid encoded hex string
			},
			"invalid hash length; should equal 20",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgPricesPrevote("abcdabcdabcdabcd", utils.TestAddress(1))
			require.NoError(t, msg.ValidateBasic())
			require.Equal(t, types.TypeMsgPricesPrevote, msg.Type())
			tc.malleate(msg)
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestMsgVote_ValidateBasic(t *testing.T) {
}

func TestMsgLsvFeederUpdate_ValidateBasic(t *testing.T) {

}

func TesMsgAddNewTicker_ValidateBasic(t *testing.T) {
}
